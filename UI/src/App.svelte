<script>
  import CreateTodoModal from "./components/createTodoModal.svelte";
  import Icon from "@iconify/svelte";
</script>

<div class="flex flex-col items-center mt-[50px]">
  <h1 class="text-4xl">Dondu</h1>
  <div>
    <button
      class="btn"
      on:click={() => document.getElementById("createTodoModal").showModal()}
      >CREATE <Icon icon="material-symbols:add" class="text-xl" /></button
    >
    <CreateTodoModal />
  </div>
</div>
