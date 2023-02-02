import type React from "react";
  
// changing from * as editor to -> {editor} affects the method of moudule import??How??
import type { editor } from "monaco-editor/esm/vs/editor/editor.api";
import MonacoEditor, { OnMount } from "@monaco-editor/react";
import { useEffect, useState } from "react";
import { editorLang } from "../atoms/monacoAtoms";
import { fileContentAtom } from '../atoms/fileReaderAtoms';
import { useAtom } from "jotai";
import Loading from "./loading";


const Editor:React.FC = () => {
  const [editor, setEditor] = useState<editor.IStandaloneCodeEditor>();
  const [lang] = useAtom(editorLang);
  const [fileContent, _] = useAtom(fileContentAtom);

  useEffect(() => {
    if (editor?.getModel()) {
      const model = editor.getModel()!;
      if (fileContent) 
        model.setValue(fileContent)
      else
        model.setValue("");
      /*
     model.pushEditOperations(
        editor.getSelections(), [
          {
            range: model.getFullModelRange(),
            text: "hello world",
          },
        ],
        () => null
      );
/
      */
    }
    

    return () => editor?.dispose()
  },[fileContent]);

const handleEditorDidMount: OnMount= (e, _) => {
    setEditor(e);
    //console.log("Monaco", n);
    //console.log("Editor", e);
  }
  return (
    <>
      <h1 className="text-slate-700 text-4xl ml-3">Editor</h1>
        <h2 className="text-lg text-extrabold text-slate-400 ml-10">
	    Langugage: {lang[0].toUpperCase() + lang.slice(1,)}
        </h2>
      <MonacoEditor
	className="w-screen h-[75vh]"
	theme = "light"
	loading= {<Loading />}
//	language={"typescript"}

	language={lang} // remount on language change -> ref hook
	options={{
	  automaticLayout: true,
	  fontSize: 18
	}}
	onMount={handleEditorDidMount}
      />

    </>
  );
}

Editor.defaultProps = {};
export default Editor;

