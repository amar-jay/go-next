import Combobox from "./combobox";
import { useState, useEffect, useRef } from "react";
import { useAtom } from "jotai";
import { fileContentAtom } from "../atoms/fileReaderAtoms";
import { setLang, langAtom } from "../atoms/monacoAtoms";
import { ExtensionType, langExt } from "../atoms/data";
const NavBar: React.FC = () => {
  const [file, setFile] = useState<FileList | null>(null);
  const trigger = useRef<FileList | null>(null);
  const [, setFileContent] = useAtom(fileContentAtom);
  const [, setlang] = useAtom(langAtom);
  useEffect(() => {
    const fetchFile = async () => {
      if (!file || !file[0]) {
	alert("Unable to open file")
	return 
      }
        file[0]
          .text()
          .then((e) => {
            setFileContent(e);
            //alert(e);
          })
          .catch((e) => alert(e));
        const ext = file[0].name.split(".").pop() as ExtensionType | undefined;
        if (ext) {
            setLang(ext, setlang);
        }
    };
    fetchFile();
  }, [file, trigger.current]);
  return (
    <header className="w-full flex items-center justify-center py-7 px-2">
      <div className="relative inline-block text-left">
        <label
          htmlFor="file-input"
          className="inline-flex w-full justify-center rounded-md bg-blue-500 px-4 py-2 font-medium text-white hover:bg-opacity-30 focus:outline-none focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75"
        >
          Open
        </label>
        <input
          id="file-input"
          onChange={(e) => {
            setFile(e.target.files);
            trigger.current = e.target.files;
          }}
          type="file"
          className="hidden"
        />
      </div>
      <div className="flex-1 text-center text-bold">
        {file && file[0].name}
      </div>
      <Combobox />
    </header>
  );
};

NavBar.defaultProps = {};

export default NavBar;
