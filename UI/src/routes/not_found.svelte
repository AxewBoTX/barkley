<script>
  import { link } from "svelte-spa-router";
</script>

<div
  class="fixed top-0 left-0 right-0 bottom-0 flex flex-col items-center justify-center gap-[20px]"
>
  <h1 class="text-3xl">The page you are looking for could not be found!</h1>
  <a class="text-2xl underline" href="/" use:link>Go Home</a>
</div>
