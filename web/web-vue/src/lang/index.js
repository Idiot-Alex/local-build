import { createI18n } from 'vue-i18n'
import zhCN from './zh-CN'
import en from './en'

const LANG_VALUE = {
  en: 'en',
  zhCN: 'zh-CN'
}

const locales = [LANG_VALUE.en, LANG_VALUE.zhCN]

const LOCAL_BUILD_LANGUAGE = 'LOCAL_BUILD_LANGUAGE'

const i18n = createI18n({
  legacy: false,
  locale: getLanguage(),
  messages: {
    'zh-CN': zhCN,
    'en': en
  }
})

export default i18n

export function getLanguage() {
  const chooseLanguage = localStorage.getItem(LOCAL_BUILD_LANGUAGE)
  if (chooseLanguage) return chooseLanguage

  // 如果没有选择语言
  const language = navigator.language.toLowerCase()
  
  for (const locale of locales) {
    if (language.indexOf(locale) > -1) {
      return locale
    }
  }
  return LANG_VALUE.zhCN
}

export function useLanguage() {

  const changeLanguage = (lang) => {
    localStorage.setItem(LOCAL_BUILD_LANGUAGE, lang)
  }

  return {
    changeLanguage,
  }
}