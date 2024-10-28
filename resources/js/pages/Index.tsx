import Layout from '@/layouts/Layout'
import { Title } from '@solidjs/meta'

export default function Welcome() {
  return (
    <>
      <Title>Welcome</Title>
      <h1>Welcome There</h1>
      <p class="text-red-500 font-bold text-3xl">Hello, welcome to your first Inertia app from Go!</p>
    </>
  )
}

Welcome.layout = Layout
