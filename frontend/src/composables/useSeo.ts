import { useHead } from '@unhead/vue'

interface SeoOptions {
  title?: string
  description?: string
  image?: string
  url?: string
}

export function useSeo(options: SeoOptions = {}) {
  const defaultTitle = 'CloudySMS - WhatsApp Business Platform'
  const defaultDescription = 'Scale your business with the open-source WhatsApp Business Platform. Real-time chat, automated workflows, and rich analytics in one powerful dashboard.'
  
  const title = options.title ? `${options.title} | CloudySMS` : defaultTitle
  const description = options.description || defaultDescription

  useHead({
    title,
    meta: [
      {
        name: 'description',
        content: description
      },
      // Open Graph
      {
        property: 'og:title',
        content: title
      },
      {
        property: 'og:description',
        content: description
      },
      {
        property: 'og:type',
        content: 'website'
      },
      // Twitter
      {
        name: 'twitter:card',
        content: 'summary_large_image'
      },
      {
        name: 'twitter:title',
        content: title
      },
      {
        name: 'twitter:description',
        content: description
      }
    ]
  })
}
