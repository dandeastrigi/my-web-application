
# E-commerce Tool

A tool to start develop an e-commerce.

----

#### About

Before start, take a look at this [image](./images/infra.png), that's structural concept of this package, the checkout can be performed from a app in webpage from browser, this application runs under *http* protocol, using *JSON data pattern* to perform data transfer, before start, add one hostname on your */etc/hosts*, supposing you use linux os, example:

```vim
127.0.0.1	localhost
127.0.0.1       myhost.com

```

----

##### Checkout

This package depends on *node*

The checkout assignments:

- Validate checkout with stripe api.
- Send checkout requests.

***Before start***
Create an account on [Stripe](https://stripe.com/en-br), and get your test keys inside of dashboard, than put your keys on *.env* file inside checkout directory.

To execute tests with this package, i used this [repo](https://github.com/tmarek-stripe/demo-react-stripe-js), provided from [Stripe](https://stripe.com/en-br).

***Starting***
To start this service, enter on *checkout* directory, and run:

```bash
npm run dev
```

----

##### Middleware-service

The *Go Middleware service* assignments:

- Receive requests from browser.
- Confirm checkout.
- Write logs (only about herself).
- Write/Read SQLite checkout database.
- Send information to be rendered on app or browser.

Before start, edit the *Dockerfile* inside folder, and put the *hostname* of you have created before in */etc/hosts*, and put your hostname on *ENV APP_ORIGIN_URL*

```vim
ENV APP_ORIGIN_URL http://myhost.com:3000
```

Or you can put * for accept all requests (use only in development)

To start service, enter on *middleware-service* directory and run:

```bash
docker-compose up --build
```

or run with go (if you like to read sqlite database is a better way)

```bash
go get ./
go run main.go
```

----

###### *Validator*

```vim
Not implemented yet
```

The *Go validation and sync* assignments:

- Check SQLite checkout database x times in hour for new checkouts
- Send information to database when get a new checkout
- Update SQLite checkout database when receive platform hook
- Write logs (only about herself)

To start service, enter on *validator* directory and run:

```bash
docker-compose up --build
```

----

###### *Administration platform*

```vim
Not implemented yet
```

The *Administration platform* assignments:

- Read/Write database.
- Admin reports/cruds.
- Hooks to service.
- Logs (only about herself).

----

###### *Backup System*

```vim
Not implemented yet
```

The *Backup system* assignments:

- Perform backup of configured databases.
