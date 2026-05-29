//ff:func feature=scan type=test control=sequence topic=express
//ff:what E2E 스캔 테스트: 선두 검증자 체인 .route().post(validate(x.y), h) → POST에 크로스파일 Joi requestBody 생성(미들웨어 선행 없음)
package express

import "testing"

func TestScan_ChainJoiLeadingValidator(t *testing.T) {
	dir := t.TempDir()

	validationSrc := `
import Joi from 'joi';
const create = {
  body: Joi.object().keys({
    name: Joi.string().required(),
    qty: Joi.number()
  })
};
export default { create };
`
	routeSrc := `
import express from 'express';
import validate from '../middlewares/validate';
import itemValidation from '../validations/item.validation';
const router = express.Router();
router.route('/items')
  .post(validate(itemValidation.create), createItem)
  .get(listItems);
export default router;
`
	appSrc := `
import express from 'express';
import itemRouter from './routes/item';
const app = express();
app.use('/items', itemRouter);
`
	writeFile(t, dir, "validations/item.validation.ts", validationSrc)
	writeFile(t, dir, "routes/item.ts", routeSrc)
	writeFile(t, dir, "app.ts", appSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	post := findEndpoint(result.Endpoints, "POST", "/items/items")
	if post == nil {
		post = findEndpoint(result.Endpoints, "POST", "/items")
	}
	if post == nil || post.Request == nil || post.Request.Body == nil {
		t.Fatal("expected POST endpoint with body from leading chain Joi validator")
	}
	if len(post.Request.Body.Fields) != 2 {
		t.Fatalf("expected 2 body fields, got %d", len(post.Request.Body.Fields))
	}
	if !fieldRequired(post.Request.Body.Fields, "name") {
		t.Error("expected name field marked required via .required()")
	}
}
