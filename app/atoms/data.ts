export type ValueOf<T> = T extends Record<any, any>? T[keyof T] : never;
export type ExtensionType = keyof typeof langExt;
export type LanguageType = ValueOf<typeof langExt>

export interface Lang {
  name: LanguageType,
  id: number, 
};

export const langExt = { 
  txt: 'PlainText',
  js: 'Javascript',
  ts: 'Typescript',
  py: 'Python',
  rs: 'Rust',
  json: 'Json',
  go: 'Go',
  astro: 'Astro',
  css: 'CSS',
  html: 'HTML',
  htm: 'HTML',
  lua: 'Lua',
  c: 'C',
  sh: 'Shell',
} as const;

export const languageList:Lang[] = [
  { id: 0, name: 'PlainText' },
  { id: 1, name: 'Javascript' },
  { id: 2, name: 'Typescript' },
  { id: 3, name: 'HTML' },
  { id: 4, name: 'Go' },
  { id: 5, name: 'Rust' },
  { id: 6, name: 'Astro' },
  { id: 7, name: 'CSS' },
  { id: 8, name: 'Lua' },
  { id: 9, name: 'C' },
  { id: 10, name: 'Shell'}
];

