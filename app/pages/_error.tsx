import Layout from "../components/layout";
const Error = (msg: unknown) => {
  // get message params and pass it as message
  return (
    <Layout>
      <h1>Error</h1>
      <p>An unknown error occurred</p>
      <p>{JSON.stringify(msg)}</p>
    </Layout>
  );
};
export default Error;
