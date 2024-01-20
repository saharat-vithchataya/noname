<script lang="ts">
  import { Calendar as CalendarIcon } from "radix-icons-svelte";
  import {
    type DateValue,
    DateFormatter,
    getLocalTimeZone
  } from "@internationalized/date";
  import { cn } from "$lib/utils";
  import { Button } from "$lib/components/ui/button";
  import { Calendar } from "$lib/components/ui/calendar";
  import * as Select from "$lib/components/ui/select";
  import * as Popover from "$lib/components/ui/popover";
  import Header from "$lib/components/Header.svelte";
  import Input from "$lib/components/ui/input/input.svelte";
	import Label from "$lib/components/ui/label/label.svelte";

  const df = new DateFormatter("en-US", {
    dateStyle: "long"
  });

  const countries = [
    { value: "th", label: "Thailand" },
    { value: "us", label: "United States" },
    { value: "jp", label: "Japan" },
    { value: "de", label: "Germany" },
    { value: "ca", label: "Canada" },
    { value: "fr", label: "France" },
    { value: "uk", label: "United Kingdom" },
    { value: "it", label: "Italy" },
    { value: "au", label: "Australia" },
    { value: "br", label: "Brazil" },
  ];

  let value: DateValue | undefined = undefined;
</script>

<Header/>
<main>
<div class="h-screen pt-20 max-w-96 mb-44 m-auto">
    <center>
      <img class="w-20 mb-10" src="https://pngfre.com/wp-content/uploads/apple-logo-6-1024x1024.png" alt="sinestrea's logo">
    </center>
    <form action="">
        <Label for="email">First name and Last name</Label>
        <Input class="bg-white  mb-1    " type="text" placeholder="First name" />
        <Input class="mb-3 bg-white" type="text" placeholder="Last name" />
        <Label for="email">Birthday and Country</Label>
      <Popover.Root>
        <Popover.Trigger asChild let:builder>
          <Button
            variant="outline"
            class={cn(
              "w-full justify-start text-left font-normal mb-1",
              !value && "text-muted-foreground"
            )}
            builders={[builder]}
          >
            <CalendarIcon class="mr-2 h-4 w-4" />
            {value ? df.format(value.toDate(getLocalTimeZone())) : "Birthday"}
          </Button>
        </Popover.Trigger>
        <Popover.Content class="w-auto p-0" align="start">
          <Calendar bind:value />
        </Popover.Content>
      </Popover.Root>
      <Select.Root>
        <Select.Trigger class="w-full mb-3">
          <Select.Value placeholder="Country" class="" />
        </Select.Trigger>
        <Select.Content>
          <Select.Group>
            <Select.Label class="">Country</Select.Label>
            {#each countries as country}
              <Select.Item value={country.value} label={country.label}
                >{country.label}</Select.Item
              >
            {/each}
          </Select.Group>
        </Select.Content>
        <Select.Input name="favoriteFruit" />
      </Select.Root>
    <Label for="email">Phone</Label>
      <Input class="mb-3 bg-white" type="phone" placeholder="Phone" />
        <Label for="email">Email and Password</Label>
      <Input class="mb-1 bg-white" type="email" placeholder="string@example.com" />
      <Input class="mb-1 bg-white" type="password" placeholder="Password" />
      <Input class="mb-3 bg-white" type="password" placeholder="Confirm Password" />
    </form>
        <Button class='w-full mb-2'>Sign up</Button>
    <a class="text-sm text-blue-400 hover:text-blue-500" href="/">Forgot your password?</a>
  </div>
</main>
