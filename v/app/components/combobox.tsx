import { Fragment } from "react";
import { Combobox, Transition } from "@headlessui/react";
import { languageList, LanguageType, Lang } from "../atoms/data";
import { langAtom } from "../atoms/monacoAtoms";
import { useAtom } from "jotai";

export default () => {
  const filteredLang = languageList.map(e => e.name);
  const [lang, setLang] = useAtom(langAtom);

  return (
    <div className="w-15 md:w-72 z-50">
      <Combobox value={lang} onChange={setLang}>
        <div className="relative mt-1">
          <div className="relative w-full cursor-default overflow-hidden rounded-lg bg-white text-left sm:text-sm">
            <Combobox.Input
              className="w-full border-none py-2 pl-3 pr-10 text-sm leading-5 bg-blue-500 text-white font-bold focus:outline-none"
              displayValue={(s:Lang["name"]) => s}
              onChange={(event:any) =>{
                if (event.target.value){
                  if (lang in filteredLang){
                    setLang(event.target.value as LanguageType)
                  }
                }
              }
              }
            />
            <Combobox.Button className="absolute inset-y-0 right-0 flex items-center pr-2 bg-blue-300 text-sm my-1 px-2 mr-1 rounded-lg">
	      Change 
            </Combobox.Button>
          </div>
          <Transition
            as={Fragment}
            leave="transition ease-in duration-100"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
            afterLeave={() => setLang(lang)}
          >
            <Combobox.Options className="absolute mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg focus:outline-none sm:text-sm">
              {filteredLang.length === 0 ? (
                <div className="relative cursor-default select-none py-2 px-4 text-red-700">
                  Nothing found.
                </div>
              ) : (
                filteredLang.map((person, k) => (
                  <Combobox.Option
                    key={k}
                    className={({ active }) =>
                      `relative cursor-default select-none py-2 text-black pl-10 pr-4 ${
                        active && "bg-blue-100 text-white"
                      }`
                    }
                    value={person}
                  >
                    {({ selected, active }) => (
                      <>
                        <span
                          className={`block truncate ${
                            selected ? "font-bold" : "font-normal"
                          }`}
                        >
                          {person}
                        </span>
                        {!!selected && (
                          <span
                            className={`absolute inset-y-0 left-0 flex items-center pl-3 ${
                              active ? "text-white" : "text-blue-600"
                            }`}
                          >
			      ✅
                          </span>
                        )}
                      </>
                    )}
                  </Combobox.Option>
                ))
              )}
            </Combobox.Options>
          </Transition>
        </div>
      </Combobox>
    </div>
  );
};
