<script lang="ts">
	import { page } from '$app/stores';
    const array: Record<string, string>[] = [
        {
            url: "/auth/register",
            name: "Inscription"
        },
        {
            url: "/auth/login",
            name: "Connexion"
        }
    ]

    let open:boolean = false;

    const currentPathname = array.find(item => {
        return item.url === $page.url.pathname
    })

</script>

<!-- <div class="w-full top-11 h-[50px] rounded-[8px] flex flex-row items-center gap-x-[6px]">
    <div class="w-[10%] rounded-[8px] h-full">
        <div class="w-full h-full flex items-center justify-center bg-white rounded-[8px]">
            <ModernArrow classes={"fill-[#191C1F] w-[18px] h-[18px] relative right-[2px]"} />
        </div>
    </div>
    <div class="w-[90%] rounded-[8px] bg-white h-full flex flex-row items-center px-2 gap-x-2">
        {#each array as item}
        <a href={item.url} class="h-[calc(100%-16px)] bg-white w-1/2 rounded-[8px] flex items-center justify-center font-semibold" class:active={item.url === $page.url.pathname}>
            <h2>{item.name}</h2>
        </a>
        {/each}
    </div>
</div> -->

<div class="flex items-center justify-center sm:p-3 fixed w-full">
    <header class="w-full md:w-[85%] h-[50px]  backdrop-blur-[2px] rounded-[8px] flex items-center  px-5">
        <div class="hidden sm:flex items-center justify-between w-full h-full">
            <h2 class="font-bold">H.</h2>
            <nav class="flex justify-center gap-x-10 font-medium text-[14px]">
                <a href="/">Accueil</a>
                <div>Authentification</div>
            </nav>
            <div></div>
        </div>
        <div class="flex sm:hidden items-center justify-between w-full px-2 h-[50%]">
            <h2 class="font-bold w-[10%]">H.</h2>
            <div class="flex justify-center w-[80%]">
                <h2 class="font-bold text-[14px]">{currentPathname?.name}</h2>
            </div>
            <div class="flex flex-col justify-between items-end w-[10%]">
                <div class="w-[11px] h-[2px] bg-black"></div>
                <div class="w-[11px] h-[2px] bg-black my-[3px]"></div>
                <div class="w-[11px] h-[2px] bg-black"></div>
            </div>
        </div>
    </header>
</div>

<style>
    .active {
        background-color: #191C1F;
        color: #fff;
    }
</style>