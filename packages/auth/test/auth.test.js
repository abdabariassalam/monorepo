process.env.NODE_ENV = 'test';

const chai = require('chai')
const chaiHttp = require('chai-http')


const app = require('../server')

const {expect} = chai

chai.use(chaiHttp)
chai.should()
describe("Test auth REST API",function(){
    describe("POST /auth/register ",function(){
        it('Should return 200 and return all user information',function(done){
            const body = {
                phone: "+6281132230455",
                name: "test1",
                role: "reporter"
            }
            chai
                .request(app)
                .post('/auth/register')
                .send(body)
                .end(function(err,res){
                    res.body.should.be.a('object')
                    expect(res.status).equal(201)
                    return err?done(err):done()
                })
        })
    })

    describe("POST /auth/login ",function(){
        it('Should return 200 and return token',function(done){
            const body = {
                phone: "+62811322334455",
                password: "GPIK"
            }
            chai
                .request(app)
                .post('/auth/login')
                .send(body)
                .end(function(err,res){
                    res.body.should.be.a('object')
                    expect(res.status).equal(200)
                    return err?done(err):done()
                })
        })
    })

    describe("GET /auth/verify ",function(){
        it('Should error 401',function(done){
            chai
                .request(app)
                .get('/auth/verify')
                .set({'x-access-token': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWRtaW4xIiwicGhvbmUiOiI4MTEyMjMzNDQ1NSIsInJvbGUiOiJhZG1pbiIsImlhdCI6MTYzNTEzMDkxOSwiZXhwIjoxNjM1MTM4MTE5fQ.sgcowLXors0Y6aHvBNPi4Fno019XNzRymUPfKUznP6Y'})
                .then((res) => {
                    expect(res).to.have.status(401)
                    const body = res.body
                    console.log(body)
                   done();
                 }).catch((err) => done(err))
        })
    })

})