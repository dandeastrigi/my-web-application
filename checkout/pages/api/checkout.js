import Stripe from "stripe";

const stripe = new Stripe(process.env.SECRET_KEY);

export default async (req, res) => {
  if (req.method === "POST") {
    try {
      const { amount } = req.body;
      
      console.log("checkout")
      res.status(200).send("checkout");
    } catch (err) {
      console.log("Erro")
      res.status(500).json({ statusCode: 500, message: err.message });
    }
  } else {
    console.log("post")
    res.setHeader("Allow", "POST");
    res.status(405).end("Method Not Allowed");
  }
};
