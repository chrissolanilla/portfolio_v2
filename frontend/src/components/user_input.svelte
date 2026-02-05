<script>
	import { onMount } from 'svelte';

  import Input from "../design-system/input.svelte";
  import Form from "../design-system/form.svelte";
  import FormButton from "../design-system/form_button.svelte";
  import Notif from "../design-system/notif.svelte";
  import DataPreview from "../components/data_preview.svelte";

  import { api } from "../utils/api.svelte.js";

  let id = $state(0);
  let name = $state("");
  let password = $state("");
  let editMode = $state(false);
  let users = $state([]);
  
  let loading = $state(false)
  let error = $state("")
  let success = $state("")

  const getUsers = async () => {
    const data = await api.get("/users");

    users = data;
  }

  const editHook = (data) => {
    if (data === null) {
      editMode = false;
      id = 0;
      name = "";
      password = "";
      return
    }
    editMode = true;
    id = data.id;
    name = data.name;
    password = data.password;
  }

  const submitForm = async () => {
    error = "";
    success = "";
    loading = true;

    const payload = {
      name,
      password
    }

    if (editMode) {
      payload.id = id;
      try {
        await api.put("/users", payload);
        name = ""
        password = ""

        success = "User has been updated successfully.";
      } catch (err) {
        error =
          err?.data?.error ||
          err?.data?.message ||
          "Failed to create user";
      } finally {
        loading = false;
      }

      return;
    }

    try {
      await api.post("/users", payload)

      name = "";
      password = "";
      success = "User created successfully!"
    } catch (err) {
      console.error(err)
      error = err || "Failed to create user"
    } finally {
      loading = false;
    }
  }

  onMount(() => {
    getUsers();
  })

</script>

<div class="flex flex-row">
  <Form submitForm={submitForm} title={"Add User"} editMode={editMode} editHook={editHook}>
    <Input label={"Name"} bind:value={name} required={true}/>
    <Input label={"Password"} bind:value={password} required={true}/>
    <FormButton loading={loading}/>
    <Notif error={error} success={success}/>
  </Form>
  <DataPreview data={users} editHook={editHook}/>
</div>

