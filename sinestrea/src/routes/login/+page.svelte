<script>
	import Header from "$lib/components/Header.svelte";
	import Button from "$lib/components/ui/button/button.svelte";
    import Input from "$lib/components/ui/input/input.svelte";


</script>
<Header/>
<main class="bg-gray-50">
    <div class=" h-screen pt-20 max-w-96 m-auto">

    <center>
        <img class="w-20 mb-10" src="https://pngfre.com/wp-content/uploads/apple-logo-6-1024x1024.png" alt="sinestrea's logo">
    </center>
    <form action="">
        <Input class="mb-1 bg-white" type="email" placeholder="example@string.com" />
        <Input class="mb-2 bg-white" type="password" placeholder="your password" />
        <Button class='w-full mb-2'>Login</Button>
    </form>
    <a class="text-sm text-blue-400 hover:text-blue-500" href="/">Forgot your password?</a>
    </div>
</main>
    
<!-- 
<div class="flex justify-center items-center h-screen ">
  <div class="w-full max-w-md p-6 bg-white rounded-md">
    <h1 class="text-3xl font-bold mb-6">Sign in to your Apple ID</h1>
    <form on:submit|preventDefault={handleLogin}>
      <Input type="email" placeholder="Apple ID" bind:value={email} class="mb-4" />
      <Input type="password" placeholder="Password" bind:value={password} class="mb-4" />
      <Button class="w-full" type="submit">Sign In</Button>
    </form>
    <div class="mt-4">
      <a href="/" class="text-sm text-blue-500">Forgot Apple ID or password?</a>
    </div>
  </div>
</div> -->