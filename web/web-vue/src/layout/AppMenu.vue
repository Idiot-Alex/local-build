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
            { label: '测试', icon: 'pi pi-fw pi-wrench', to: '/env/crud' },
        ]
    },
    {
        label: 'Pages',
        icon: 'pi pi-fw pi-briefcase',
        to: '/pages',
        items: [
            {
                label: 'Landing',
                icon: 'pi pi-fw pi-globe',
                to: '/landing'
            },
            {
                label: 'Auth',
                icon: 'pi pi-fw pi-user',
                items: [
                    {
                        label: 'Login',
                        icon: 'pi pi-fw pi-sign-in',
                        to: '/auth/login'
                    },
                    {
                        label: 'Error',
                        icon: 'pi pi-fw pi-times-circle',
                        to: '/auth/error'
                    }
                ]
            },
            {
                label: 'Crud',
                icon: 'pi pi-fw pi-pencil',
                to: '/pages/crud'
            },
            {
                label: 'Timeline',
                icon: 'pi pi-fw pi-calendar',
                to: '/pages/timeline'
            }
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
