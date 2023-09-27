<script lang="ts">
    import {page} from "$app/stores";

    console.log($page)

    const navigation = ["register", "login"];

    const url: string = $page.url.href.split("/").at(-1) as string;

    const other_url_segment = navigation.find((u => u !== url))
</script>

    <a class="font-medium text-sm text-center" href={`http://localhost:5173/auth/${other_url_segment}`}>
        {#if url === "register"}
        Déjà inscrit ? <span class="font-medium text-accent-violet-d"> Connectez-vous ici</span>
        {:else}
        Pas encore de compte ?
        {/if}
    </a>
