<script>
  import Router from "svelte-spa-router";
  import Home from "./routes/home.svelte";
  import About from "./routes/about.svelte";
  import Posts from "./routes/posts.svelte";
  import NotFound from "./routes/not_found.svelte";

  const routes = {
    "/": Home,
    "/about": About,
    "/posts": Posts,
    "*": NotFound,
  };
</script>

<body>
  <Router {routes} />
</body>
