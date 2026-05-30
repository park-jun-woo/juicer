//ff:type feature=scan type=test topic=hono
//ff:what scanInternalsSrc 테스트 보조 선언
package hono

const scanInternalsSrc = `
import { Hono } from "hono"
const app = new Hono()
app.get("/users/:id", (c) => c.json({ id: 1 }))
app.post("/users", (c) => c.json({ ok: true }, 201))
`
