<script>
  let { submitForm, title, children, editMode, editHook } = $props();
</script>

<form
	onsubmit={submitForm}
	class="min-w-xl max-w-xl space-y-6 rounded-2xl border border-gray-200 bg-white p-8 shadow-sm"
>
  
  {#if editMode}
    <div>
      <a class="text-fuchsia-950"
        onclick={() => editHook(null)}
      >
        you are now in edit mode. i love you. click this to exit edit mode
      </a>
    </div>
  {/if}
	<h2 class="text-xl font-semibold text-gray-900">
    {title}
	</h2>

  {@render children()}

</form>
