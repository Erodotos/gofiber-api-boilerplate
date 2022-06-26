# gofiber-api-boilerplate

This is a very basic boilerplate API template I created using the [go-fiber](https://gofiber.io/) web framework. To be honest this framework is pretty fast!

In this get-started project we are interested on the book business (so unique 😂). We implement the following end-points: 

- POST   : /api/book/  
- GET    : /api/book/:id
- GET    : /api/book/all
- PUT    : /api/book/:id
- DELETE : /api/book/:id

**Note:** The update end-point is not workin, due to an incosistency between [GORM](https://gorm.io/index.html) and MySql server. Old version seem to be buggy.

