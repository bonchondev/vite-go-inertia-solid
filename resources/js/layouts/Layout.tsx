import { Link } from 'inertia-adapter-solid'
import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/solid-query'

const queryClient = new QueryClient();

export default function Layout(props) {
  return (
    <QueryClientProvider client={queryClient}>
      <main>
        <header class="flex flex-row p-3 gap-x-5">
          <Link href="/" class="text-5xl font-bold underline">Home</Link>
          <Link href="/about" class="text-5xl font-bold underline">About</Link>
        </header>
        <article class='p-3'>{props.children}</article>
      </main>
    </QueryClientProvider>
  )
}
