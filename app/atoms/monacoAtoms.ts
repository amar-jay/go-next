import { atom } from 'jotai'
import type {ExtensionType, LanguageType} from './data'
import { languageList, langExt } from './data';


// Create your atoms and derivatives
export const langAtom = atom(languageList[0].name);
//const selectedLangAtom = atom(languageList[0].name);
export const editorLang = atom((get) => get(langAtom).toLowerCase());
export const filteredLangList = languageList.map(e => e.name);

export const setLang = (ext: ExtensionType| undefined, setLang: (x:LanguageType) => void) => {
  if (!ext || ext.length < 1 || !langExt[ext]){
    return new Error(`Extension "${ext}" does not exist`);
  }
  const lang = langExt[ext];
  setLang(lang);
}


