import Editor from "../components/editor"
import Layout from "../components/layout"
import Navbar from "../components/navbar"

export default function EditorPage() {
  return (
    <Layout>
      <Navbar />
      <h1 className="text-blue-500">Editor</h1>
      <Editor />
    </Layout>
  )
}
