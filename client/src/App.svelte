<script lang="ts">
  async function getAccounts() {
   let response = await fetch("http://localhost:3000/api/accounts")
   const data = response.json()
   return data
  }
</script>


<main>
  <div class="card">
    {#await getAccounts()}
      loading accounts...
    {:then account}
      {account[0].full_name}
    {:catch someError}
      System error: {someError.message}.
    {/await}
   
  </div>

</main>