<script>
  import { Link } from '@lucide/svelte';
  import { Newspaper } from '@lucide/svelte';

  let { project } = $props();

  const projectBgClass = {
    work: 'bg-work-project',
    personal: 'bg-personal-project',
  }[project.type] ?? 'bg-gray-200';

</script>
<div
  class={`relative flex flex-col items-center p-8 w-full min-w-[280px] max-w-md mx-auto
          rounded-xl border-4 border-black
          after:content-[''] after:absolute after:inset-0
          after:rounded-xl after:bg-black
          after:translate-x-3 after:translate-y-3
          after:-z-10
          ${projectBgClass}`}
>
  <img
    src={project.image}
    alt={project.name}
    class="block max-w-full h-auto rounded-lg border-2 border-black"
  />
  <h3 class="text-lg font-semibold mt-4 text-gray-900 leading-tight">
    {project.name}
  </h3>
  <p class="text-center mt-2 text-sm leading-relaxed">
    {project.description}
  </p>
  <div class="flex w-full justify-center gap-4 mt-6">
    <a href={project.githubLink} class="p-3 rounded-full bg-gray-100 border-2 border-black hover:cursor-pointer" target="_blank">
      <Link />
    </a>
    {#if project.blogLink}
      <a href={project.blogLink} class="p-3 rounded-full bg-gray-100 border-2 border-black">
        <Newspaper />
      </a>
    {/if}
  </div>
</div>
