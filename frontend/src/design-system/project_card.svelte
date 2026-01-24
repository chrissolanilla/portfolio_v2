<script>
  let { projectType } = $props();

  const projectBgClass = {
    work: 'bg-work-project',
    personal: 'bg-personal-project',
  }[projectType] ?? 'bg-gray-200';

</script>

<div
  class={`relative flex flex-col items-center p-8 max-w-md w-full mx-auto
          rounded-xl border-4 border-black
          after:content-[''] after:absolute after:inset-0
          after:rounded-xl after:bg-black
          after:translate-x-3 after:translate-y-3
          after:-z-10
          ${projectBgClass}`}
>

  <div class="w-full h-40 bg-gray-100 rounded-lg border-2 border-black"></div>

  <h3 class="text-lg font-semibold mt-4 text-gray-900">
    Project Name
  </h3>

  <p class="text-center mt-2 text-sm">
    This is the project's description lorem ipsum dolor
  </p>

  <div class="flex w-full justify-center gap-4 mt-6">
    <div class="p-3 rounded-full bg-gray-100 border-2 border-black">
      Icon 1
    </div>

    <div class="p-3 rounded-full bg-gray-100 border-2 border-black">
      Icon 2
    </div>
  </div>
</div>
