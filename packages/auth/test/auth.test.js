const chai = require('chai')
const chaiHttp = require('chai-http')

const app = require('../server')

const {expect} = chai

chai.use(chaiHttp)
chai.should()
describe("Test Messages REST API",function(){
    describe("GET /messages ",function(){
        it('Should return 200 and return all messages',function(done){
            chai
                .request(app)
                .get('/messages')
                .end(function(err,res){
                    res.body.should.be.a('object')
                    expect(res.status).equal(200)
                    return err?done(err):done()
                })
        })
    })

    describe("POST /messages", function(){
        it('Should send messages from request', function(done){
            const msg_text = {
                message: 'hellow its me'
            }
            chai.request(app)
                .post('/messages')
                .send(msg_text)
                .end(function(err,res){
                    expect(res.status).equal(201)
                    expect(res.body.message).equal("Sent")
                    return err?done(err):done()
                })
        })
    })

})