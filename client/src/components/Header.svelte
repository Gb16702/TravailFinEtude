<script lang="ts">
    import Logo from "./Logo.svelte";


</script>


<div class="w-full h-[65px] flex justify-center border-b border-black bg-white">
<div class="w-[95%] max-w-[1500px] h-full flex items-center justify-between px-4">
    <div class="w-1/3">
        <div>
            <Logo />
        </div>
    </div>
    <div class="w-1/3">
        <div class="flex items-center justify-center gap-x-3">
            <h2>Test</h2>
            <h2>Test</h2>
            <h2>Test</h2>
            <h2>Test</h2>
            <h2>Test</h2>
        </div>
    </div>
        <div class="w-1/3">
            <div class="flex flex-row items-center justify-end gap-x-2">
            <a href="/auth">
                Authentification
            </a>
            </div>
        </div>
    </div>
</div>
