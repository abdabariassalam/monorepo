const express = require('express');
const router = express.Router();
var low = require('lowdb');
var FileSync = require('lowdb/adapters/FileSync');
var adapter = new FileSync('./../db/db.json');
var db = low(adapter);
const jwt = require("jsonwebtoken");
var generator = require('generate-password');

// registrasi user
router.post('/register', (req, res) => {
  // 
  const { phone, name, role } = req.body;

  // Validate user input
  if (!(phone && name && role)) {
    res.status(400).send("All input is required");
  }

  const isPhone = NormalisasiPhone(phone)
  if (!isPhone){
    res.status(400).send("Your Number not valid");
    return
  }
  // check if user already exist
  // Validate if user exist in our database
  const oldUser = getUser(isPhone) 

  if (oldUser) {
    return res.status(409).send("User Already Exist. Please Login");
  }

  //Encrypt user password
  var isSamePass = false
  var generatePassword
  while (!isSamePass) {
    generatePassword = generator.generate({
      length: 4,
      numbers: true
    });
  
    isSamePass = checkSimilarPassword(generatePassword) 
    if(!isSamePass){
      isSamePass = true 
    }
  }
  
  // dump data to database
  user = db.get("users").push({
    phone: isPhone,
    name: name,
    role: role, 
    password: generatePassword,
  }).write();

  // return new user
  res.status(201).json(getUser(isPhone));
});

// api untuk login menggunakan phone dan password
router.post('/login', (req, res) => {
    // Get user input
    const { phone, password } = req.body;

    // Validasi user input
    if (!(phone && password)) {
      res.status(400).send("All input is required");
      return
    }

    // standarisasi nomor ponsel
    const isPhone = NormalisasiPhone(phone)
    if (!isPhone){
      res.status(400).send("Your Number not valid");
      return
    }

    // getUser from database
    var user = getUser(isPhone) 

    // check if user exist and password match
    if (user && password == user.password) {
      // Create token
      const token = jwt.sign(
        { name: user.name, phone: user.phone, role: user.role },
        process.env.JWT_SECRET,
        {
          expiresIn: "2h", // set expired token in 2 jam
        }
      );

      // save user token
      user.token = token;

      res.status(200).json(user);
    }else{
      res.status(400).send("Wrong phone or Password");
    }

});

router.get('/verify', (req, res) => {

  const token = req.headers["x-access-token"];
  if (!token) {
    return res.status(403).send("A token is required");
  }

  try {
    const decoded = jwt.verify(token, process.env.JWT_SECRET);
    user = decoded;
  } catch (err) {
    return res.status(401).send("Bad Token");
  }

  res.status(200).json(user);
});

// get user di database
function getUser(phone){
    db.read(); // read data di jsondb
    var putusers=db.get('users'); // use users as a section di jsondb
    var user = putusers.find({ phone: phone }).value() // masukkan value ke variable user
    return user
}

// check similar password
function checkSimilarPassword(password){
  db.read(); // read data di jsondb
  var putpass=db.get('password'); // use password as a section di jsondb
  var pass = putpass.find({ password: password }).value() // masukkan value ke variable pass
  return pass
}

// standarisasi dan validasi nomor hp
function NormalisasiPhone(phone){
  phone = String(phone).trim();
  if(phone.startsWith('+62')){
    phone = phone.slice(3);
  } else if (phone.startsWith('62')){
    phone = phone.slice(2);
  } else {return}
  return phone;
}

module.exports = router;