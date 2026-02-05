<script>
  let { label, required, value = $bindable() } = $props();
</script>


<div>
  <label class="block text-sm font-medium text-gray-700">
    {label} {#if required}<span class="text-red-500">*</span>{/if}
  </label>
  <input
    type="text"
    bind:value={value}
    required={required}
    class="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 text-sm shadow-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
  />
</div>
