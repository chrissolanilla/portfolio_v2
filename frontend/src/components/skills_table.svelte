<script>
  let { skills } = $props();

  let categories = $derived.by(() => {
    return [...new Set(skills.map(skill => skill.category))];
  });

  let selectedCategory = $derived.by(() => {
    if (categories.length > 0) {
      return categories[0]
    }
  })

  let filteredSkills = $derived.by(() => {
    return skills.filter(skill => skill.category === selectedCategory);
  });
</script>

<section class="flex flex-wrap justify-center p-1">
<div class="w-full overflow-x-auto">
  <table class="border-collapse border border-black max-w-7xl min-w-full bg-skills">
    <thead class="bg-gray-100 text-center">
      <tr>
        <th class="border border-black px-4 py-2 text-center">
          <p>Skills:</p>
        </th>
        <th class="border border-black px-4 py-2 text-center">
          <label for="categories">Category:</label>
          <select
            name="categories"
            id="categories"
            class="text-center"
            bind:value={selectedCategory}
          >
            {#each categories as category}
              <option value={category}>{category}</option>
            {/each}
          </select>
        </th>
      </tr>
    </thead>
    <tbody class="text-center">
      {#each filteredSkills as skill}
        <tr>
          <td class="border border-black px-4 py-2">{skill.name}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
</section>

