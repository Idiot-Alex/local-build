<script setup>
import { ref } from 'vue'
import Dock from 'primevue/dock';


defineProps({
  msg: String,
})

const count = ref(0)
const items = ref([
    {
        label: 'Finder',
        icon: "https://primefaces.org/cdn/primevue//images/dock/finder.svg",
        command: () => {
            displayFinder.value = true;
        }
    },
    {
        label: 'Terminal',
        icon: "https://primefaces.org/cdn/primevue//images/dock/terminal.svg",
        command: () => {
            displayTerminal.value = true;
        }
    },
    {
        label: 'App Store',
        icon: "https://primefaces.org/cdn/primevue//images/dock/appstore.svg",
        command: () => {
            toast.add({ severity: 'error', summary: 'An unexpected error occurred while signing in.', detail: 'UNTRUSTED_CERT_TITLE', group: 'tc', life: 3000 });
        }
    },
    {
        label: 'Safari',
        icon: "https://primefaces.org/cdn/primevue//images/dock/safari.svg",
        command: () => {
            toast.add({ severity: 'warn', summary: 'Safari has stopped working', group: 'tc', life: 3000 });
        }
    },
    {
        label: 'Photos',
        icon: "https://primefaces.org/cdn/primevue//images/dock/photos.svg",
        command: () => {
            displayPhotos.value = true;
        }
    },
    {
        label: 'GitHub',
        icon: "https://primefaces.org/cdn/primevue//images/dock/github.svg",
    },
    {
        label: 'Trash',
        icon: "https://primefaces.org/cdn/primevue//images/dock/trash.png",
        command: () => {
            toast.add({ severity: 'info', summary: 'Empty Trash', life: 3000 });
        }
    }
]);
</script>

<template>
  <Dock :model="items">
    <template #item="{ item }">
        <a v-tooltip.top="item.label" href="#" class="p-dock-link" @click="onDockItemClick($event, item)">
            <img :alt="item.label" :src="item.icon" style="width: 100%" />
        </a>
    </template>
  </Dock>
  <h1>{{ msg }}</h1>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
