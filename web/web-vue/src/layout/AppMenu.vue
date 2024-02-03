<script setup>
import { ref } from 'vue';

import AppMenuItem from './AppMenuItem.vue';

const model = ref([
    {
        label: '主页',
        items: [{ label: '仪表盘', icon: 'pi pi-fw pi-home', to: '/' }]
    },
    {
        label: '项目及构建',
        items: [
            { label: '项目管理', icon: 'pi pi-fw pi-bars', to: '/project' },
            { label: '构建计划', icon: 'pi pi-fw pi-eye', to: '/build-plan' },
        ]
    },
    {
        label: '环境',
        items: [
            { label: '构建工具', icon: 'pi pi-fw pi-wrench', to: '/tool' },
        ]
    },
]);
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
