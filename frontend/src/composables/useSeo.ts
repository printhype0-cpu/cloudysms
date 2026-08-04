import { useHead } from '@unhead/vue'

interface SeoOptions {
  title?: string
  description?: string
  image?: string
  url?: string
  canonical?: string
  schemas?: any[]
}

export function useSeo(options: SeoOptions = {}) {
  const defaultTitle = 'Turn WhatsApp into Your #1 Sales Channel | CloudySMS'
  const defaultDescription = 'CloudySMS - The Modern, Open-Source WhatsApp Business Platform. Automate conversations, build chatbots, and engage your customers at scale.'
  const defaultUrl = 'https://www.cloudysms.com'
  
  const title = options.title ? `${options.title} | CloudySMS` : defaultTitle
  const description = options.description || defaultDescription
  const url = options.url || defaultUrl
  const canonical = options.canonical || url

  const metaData: any = [
    { name: 'description', content: description },
    // Open Graph
    { property: 'og:title', content: title },
    { property: 'og:description', content: description },
    { property: 'og:type', content: 'website' },
    { property: 'og:url', content: url },
    // Twitter
    { name: 'twitter:card', content: 'summary_large_image' },
    { name: 'twitter:title', content: title },
    { name: 'twitter:description', content: description }
  ]

  if (options.image) {
    metaData.push({ property: 'og:image', content: options.image })
    metaData.push({ name: 'twitter:image', content: options.image })
  }

  const scripts = options.schemas ? options.schemas.map(schema => ({
    type: 'application/ld+json',
    innerHTML: JSON.stringify(schema)
  })) : []

  useHead({
    title,
    link: [
      { rel: 'canonical', href: canonical }
    ],
    meta: metaData,
    script: scripts
  })
}
