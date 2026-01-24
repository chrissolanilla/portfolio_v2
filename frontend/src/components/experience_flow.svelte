<script>
  export let items = [];
</script>

<div class="flex flex-wrap items-center justify-center gap-6 py-16">
  {#each items as item, index}
    <div class="relative">
      <div
        class="min-w-[220px] rounded-2xl border-4 border-black bg-experience p-5 text-center
          after:content-[''] after:absolute after:inset-0
          after:rounded-xl after:bg-black
          after:translate-x-3 after:translate-y-3
          after:-z-10
        "
      >
        <div class="text-lg font-semibold">
          {item.title}
        </div>

        {#if item.subtitle}
          <div class="mt-1 text-sm">
            {item.subtitle}
          </div>
        {/if}

        {#if item.description}
          <div class="mt-2 text-sm">
            {item.description}
          </div>
        {/if}
      </div>
    </div>

    {#if index < items.length - 1}
      <div class="select-none text-3xl">
        →
      </div>
    {/if}
  {/each}
</div>
