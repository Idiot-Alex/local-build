<script setup>
import { ref } from 'vue'
import AppMenuItem from './AppMenuItem.vue'


const model = ref([
    {
        label: 'home',
        items: [{ label: 'dashboard', icon: 'pi pi-fw pi-home', to: '/' }]
    },
    {
        label: 'project_build',
        items: [
            { label: 'project_manage', icon: 'pi pi-fw pi-bars', to: '/project' },
            { label: 'build_plan', icon: 'pi pi-fw pi-eye', to: '/build-plan' },
        ]
    },
    {
        label: 'env',
        items: [
            { label: 'build_tool', icon: 'pi pi-fw pi-wrench', to: '/tool' },
        ]
    },
])
</script>

<template>
    <ul class="layout-menu">
        <template v-for="(item, i) in model" :key="item">
            <app-menu-item v-if="!item.separator" :item="item" :index="i"></app-menu-item>
            <li v-if="item.separator" class="menu-separator"></li>
        </template>
    </ul>
</template>

<style lang="scss" scoped></style>
