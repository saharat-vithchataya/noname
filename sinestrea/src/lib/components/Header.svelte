<script lang="ts">
    import {Button} from "$lib/components/ui/button"
</script>
<nav class="flex justify-between items-center px-4 py-4 border-b">
    <div class="flex items-center text-center space-x-4">
        <a href="/" class="text-2xl font-medium">Sinestrea</a>
        <a href="/" class="text-sm font-medium text-muted-foreground transition-colors hover:text-primary">Home</a>
        <a href="/" class="text-sm font-medium text-muted-foreground transition-colors hover:text-primary">Discover</a>
        <a href="/planning" class="text-sm font-medium text-muted-foreground transition-colors hover:text-primary">Planing</a>
        <a href="/" class="text-sm font-medium text-muted-foreground transition-colors hover:text-primary">Settings</a>
    </div>
    <div>
        <a href="/signup"><Button>Sign up</Button></a>
        <a href="/login"><Button variant="ghost">Login </Button></a>
    </div>
</nav>