import { Fragment } from "react";
import { Combobox, Transition } from "@headlessui/react";
import { languageList, LanguageType, Lang } from "../atoms/data";
import { langAtom } from "../atoms/monacoAtoms";
import { useAtom } from "jotai";

export default () => {
  const filteredLang = languageList.map(e => e.name);
  const [lang, setLang] = useAtom(langAtom);

  return (
    <div className="w-72 z-50">
      <Combobox value={lang} onChange={setLang}>
        <div className="relative mt-1">
          <div className="relative w-full cursor-default overflow-hidden rounded-lg bg-white text-left shadow-md focus:outline-none focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75 focus-visible:ring-offset-2 focus-visible:ring-offset-teal-300 sm:text-sm">
            <Combobox.Input
              className="w-full border-none py-2 pl-3 pr-10 text-sm leading-5 text-gray-900 focus:ring-0"
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
            <Combobox.Button className="absolute inset-y-0 right-0 flex items-center pr-2">
	      hhh
            </Combobox.Button>
          </div>
          <Transition
            as={Fragment}
            leave="transition ease-in duration-100"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
            afterLeave={() => setLang(lang)}
          >
            <Combobox.Options className="absolute mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
              {filteredLang.length === 0 ? (
                <div className="relative cursor-default select-none py-2 px-4 text-gray-700">
                  /Nothing found.
                </div>
              ) : (
                filteredLang.map((person, k) => (
                  <Combobox.Option
                    key={k}
                    className={({ active }) =>
                      `relative cursor-default select-none py-2 pl-10 pr-4 ${
                        active ? "bg-teal-600 text-white" : "text-gray-900"
                      }`
                    }
                    value={person}
                  >
                    {({ selected, active }) => (
                      <>
                        <span
                          className={`block truncate ${
                            selected ? "font-medium" : "font-normal"
                          }`}
                        >
                          {person}
                        </span>
                        {!!selected && (
                          <span
                            className={`absolute inset-y-0 left-0 flex items-center pl-3 ${
                              active ? "text-white" : "text-teal-600"
                            }`}
                          >
			      xxx
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
