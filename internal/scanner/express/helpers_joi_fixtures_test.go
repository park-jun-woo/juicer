//ff:type feature=scan type=model topic=express
//ff:what 테스트 픽스처: Joi 크로스파일 검증 스캔용 소스 문자열
package express

const joiValidationFixture = `
import Joi from 'joi';
const register = {
  body: Joi.object().keys({
    email: Joi.string().required().email(),
    password: Joi.string().required()
  })
};
const verifyEmail = {
  query: Joi.object().keys({
    token: Joi.string().required()
  })
};
export default { register, verifyEmail };
`

const joiRouteFixture = `
import express from 'express';
import validate from '../middlewares/validate';
import authValidation from '../validations/auth.validation';
const router = express.Router();
router.post('/register', validate(authValidation.register), register);
router.post('/verify-email', validate(authValidation.verifyEmail), verifyEmail);
export default router;
`

const joiAppFixture = `
import express from 'express';
import authRouter from './routes/auth';
const app = express();
app.use('/auth', authRouter);
`
