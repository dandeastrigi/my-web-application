function onGooglePayLoaded() {
    window.googlePayClient = new google.payments.api.PaymentsClient({
      environment: "TEST"
    });
}

const { googlePayClient } = window;

const baseCardPaymentMethod = {
  type: "CARD",
  parameters: {
    allowedCardNetworks: ["VISA", "MASTERCARD"],
    allowedAuthMethods: ["PAN_ONLY", "CRYPTOGRAM_3DS"]
  }
};

const googlePayBaseConfiguration = {
    apiVersion: 2,
    apiVersionMinor: 0,
    allowedPaymentMethods: [baseCardPaymentMethod]
  };