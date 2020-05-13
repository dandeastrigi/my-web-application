import FormField from "./FormField";

const BillingDetailsFields = () => {
  return (
    <>
      <FormField
        name="name"
        label="Nome"
        type="text"
        placeholder="Joel silva"
        required
      />
      <FormField
        name="email"
        label="Email"
        type="email"
        placeholder="joel.silva@gmail.com"
        required
      />
      <FormField
        name="address"
        label="Endereço"
        type="text"
        placeholder="Av. Rendeiras, 667"
        required
      />
      <FormField
        name="city"
        label="Cidate"
        type="text"
        placeholder="Florianópolis"
        required
      />
      <FormField
        name="state"
        label="Estado"
        type="text"
        placeholder="Santa Catarina"
        required
      />
      <FormField
        name="zip"
        label="CEP"
        type="text"
        placeholder="88065100"
        required
      />
    </>
  );
};

export default BillingDetailsFields;
