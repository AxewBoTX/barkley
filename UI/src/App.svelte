<script>
  import toast, { Toaster } from "svelte-french-toast";
  import Icon from "@iconify/svelte";

  const fetchTodos = async () => {
    const res = await fetch("/api/todo");
    const data = await res.json();
    return data;
  };
</script>

<div class="flex flex-col items-center mt-[50px] px-4">
  <h1 class="text-5xl mb-[45px]">Dondu</h1>
  <!-- Create Todo Form -->
  <form
    class="join w-full max-w-[800px] flex items-center justify-end"
    on:submit|preventDefault={(e) => {
      const form_data = new FormData(e.currentTarget);
      const formData = {};
      for (let i of form_data) {
        const [key, value] = i;
        formData[key] = value;
      }
      if (formData.title.trim().length == 0 || formData.title.trim() == "") {
        toast.error("Input is needed!", {
          position: "top-right",
        });
      }
      // ------> Handle Data Posting
    }}
  >
    <input
      name="title"
      class="input input-bordered join-item w-full max-w-[200px]"
      placeholder="Title"
    />
    <button class="btn btn-primary btn)-outline join-item" type="submit"
      ><Icon icon="material-symbols:add" class="text-2xl" /></button
    >
  </form>
  <!-- Todo List -->
  <div>
    {#await fetchTodos()}
      <div class="text-2xl">Loading...</div>
    {:then data}
      <div class="">
        {data.length}
      </div>
    {/await}
  </div>
</div>
<Toaster />
