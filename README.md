### E-commerce Tool
A tool to start develop an e-commerce.

----
#### About
Before start, take a look at this [image](./images/infra.png), that's a structural part of solution, the checkout can be performed from a app in webpage from browser, or from an app, that's because of *Go Middleware service*  work with *http* protocol, using *JSON data pattern* to perform data transfer, after this add a hostname on your */etc/hosts*, example:

```
127.0.0.1	localhost
127.0.0.1       myhost.com

```

---
###### *checkout* path
This package depends on *node*

The assignments of checkt is:
- Validate checkout with stripe api
- Send checkout requests

To execute tests in my branch, i used this [repo](https://github.com/tmarek-stripe/demo-react-stripe-js), i just add a request to my service when checkout is confirmed by Stripe.

To start this service, enter on *checkout* directory, and run:
```
npm run dev
```

---
###### *Go Middleware service*
The *Go Middleware service* assignments:
- Receive requets from browser
- Confirm checkout
- Write logs (only about herself)
- Write/Read SQLite checkout database
- Send information to be rendered on app or browser

Before start, edit the *Dockerfile* inside folder, and put the *hostname* of you have created before in */etc/hosts*, and put your hostname on *ENV APP_ORIGIN_URL*
```
ENV APP_ORIGIN_URL http://myhost.com:3000
```
To start service, enter on *middleware-service* directory and run:
```
docker-compose up --build
```

---
###### *Go validation and sync service*
The *Go validation and sync* assignments:
- Check SQLite checkout database x times in hour for new checkouts
- Send information to database when get a new checkout
- Update SQLite checkout database when receive platform hook
- Write logs (only about herself)

To start service, enter on *validator* directory and run:
```
docker-compose up --build
```
---
###### *Administration platform*
```
Not implemented yet
```

The *Administration platform* assignments:
- Read/Write database
- Admin reports/cruds
- Hooks to service
- Logs (only about herself)

---
###### *Backup System*
```
Not implemented yet
```
The *Backup system* assignments:
- Perform backup of configured databases