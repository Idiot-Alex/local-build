<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue';
import { useLayout } from '@/layout/composables/layout';
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n'
import { useLanguage } from '@/lang/index'

const { layoutConfig, onMenuToggle, isDarkTheme, toggleTheme } = useLayout();

const outsideClickListener = ref(null);
const topbarMenuActive = ref(false);
const router = useRouter();

const { t, locale } = useI18n()
const { changeLanguage } = useLanguage()

const language_menu = ref()
const items = ref([
    {
        label: '中文',
        icon: 'pi pi-file',
        command: () => {
            locale.value = 'zh-CN'
            changeLanguage(locale.value)
        }
    },
    {
        separator: true
    },
    {
        label: 'English',
        icon: 'pi pi-file-edit',
        command: () => {
            locale.value = 'en'
            changeLanguage(locale.value)
        }
    },
])

const toggleLanguage = (event) => {
    language_menu.value.toggle(event)
};

onMounted(() => {
    bindOutsideClickListener();
});

onBeforeUnmount(() => {
    unbindOutsideClickListener();
});

const toggleDark = () => {
    let theme = layoutConfig.theme.value
    let mode = 'light'
    // current is dark
    if (layoutConfig.darkTheme.value) {
        theme = theme.replace('dark', 'light')
        mode = 'light'
    } else {
        theme = theme.replace('light', 'dark')
        mode = 'dark'
    }
    toggleTheme(theme, mode)
}

const logoUrl = computed(() => {
    return `layout/images/${layoutConfig.darkTheme.value ? 'logo-white' : 'logo-dark'}.svg`
})

const onTopBarMenuButton = () => {
    topbarMenuActive.value = !topbarMenuActive.value;
};
const onSettingsClick = () => {
    topbarMenuActive.value = false;
    router.push('/documentation');
};
const topbarMenuClasses = computed(() => {
    return {
        'layout-topbar-menu-mobile-active': topbarMenuActive.value
    };
});

const bindOutsideClickListener = () => {
    if (!outsideClickListener.value) {
        outsideClickListener.value = (event) => {
            if (isOutsideClicked(event)) {
                topbarMenuActive.value = false;
            }
        };
        document.addEventListener('click', outsideClickListener.value);
    }
};
const unbindOutsideClickListener = () => {
    if (outsideClickListener.value) {
        document.removeEventListener('click', outsideClickListener);
        outsideClickListener.value = null;
    }
};
const isOutsideClicked = (event) => {
    if (!topbarMenuActive.value) return;

    const sidebarEl = document.querySelector('.layout-topbar-menu');
    const topbarEl = document.querySelector('.layout-topbar-menu-button');

    return !(sidebarEl.isSameNode(event.target) || sidebarEl.contains(event.target) || topbarEl.isSameNode(event.target) || topbarEl.contains(event.target));
};
</script>

<template>
    <div class="layout-topbar">
        <router-link to="/" class="layout-topbar-logo">
            <img :src="logoUrl" alt="logo" />
            <span></span>
        </router-link>

        <button class="p-link layout-menu-button layout-topbar-button" @click="onMenuToggle()">
            <i class="pi pi-bars"></i>
        </button>

        <button class="p-link layout-topbar-menu-button layout-topbar-button" @click="onTopBarMenuButton()">
            <i class="pi pi-ellipsis-v"></i>
        </button>

        <div class="layout-topbar-menu" :class="topbarMenuClasses">
            <button @click="toggleDark()" class="p-link layout-topbar-button">
                <i class="pi" :class="isDarkTheme ? 'pi-moon' : 'pi-sun'"></i>
                <span>{{ isDarkTheme ? 'Moon' : 'Sun' }}</span>
            </button>
            <button @click="toggleLanguage" class="p-link layout-topbar-button" aria-haspopup="true" aria-controls="language-menu">
                <i class="pi pi-language"></i>
                <span>Language</span>
            </button>
            <TieredMenu ref="language_menu" id="language-menu" :model="items" popup w-100px />
            <button @click="onSettingsClick()" class="p-link layout-topbar-button">
                <i class="pi pi-cog"></i>
                <span>Settings</span>
            </button>
        </div>
    </div>
</template>

<style lang="scss" scoped></style>
