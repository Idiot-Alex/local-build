<script setup>
import { FilterMatchMode } from 'primevue/api'
import { ref, onMounted, onBeforeMount } from 'vue'
import { useToast } from 'primevue/usetoast'
import { toolList, saveTool } from '@/api/tool.js'

const toast = useToast()

const dataList = ref(null)
const editDialog = ref(false)
const deleteDialog = ref(false)
const deleteDialogs = ref(false)
const tempData = ref({})
const selectedProducts = ref(null)
const dt = ref(null)
const filters = ref({})
const submitted = ref(false)
const types = ref([
    { label: 'GIT', value: 'GIT' },
    { label: 'JDK', value: 'JDK' },
    { label: 'MAVEN', value: 'MAVEN' },
    { label: 'NODE', value: 'NODE' },
])

onBeforeMount(() => {
    initFilters()
})
onMounted(() => {
    loadList()
})
const loadList = () => {
    toolList({}).then(res => {
        dataList.value = res.data
    })
}
const formatCurrency = (value) => {
    return value.toLocaleString('en-US', { style: 'currency', currency: 'USD' })
}

const openNew = () => {
    tempData.value = {}
    submitted.value = false
    editDialog.value = true
}

const hideDialog = () => {
    editDialog.value = false
    submitted.value = false
}

const saveData = () => {
    submitted.value = true
    saveTool(tempData.value).then(res => {
        console.log(res)
    })
    if (tempData.value.Name && tempData.value.Name.trim() && tempData.value.price) {
        if (tempData.value.id) {
            tempData.value.type = tempData.value.type.value ? tempData.value.type.value : tempData.value.type;
            products.value[findIndexById(tempData.value.ID)] = tempData.value;
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Updated', life: 3000 });
        } else {
            tempData.value.ID = createId();
            tempData.value.code = createId();
            tempData.value.image = 'product-placeholder.svg';
            tempData.value.type = tempData.value.type ? tempData.value.type.value : 'INSTOCK';
            products.value.push(tempData.value);
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Created', life: 3000 });
        }
        editDialog.value = false;
        tempData.value = {};
    }
};

const editProduct = (editProduct) => {
    tempData.value = { ...editProduct }
    console.log(tempData)
    editDialog.value = true
};

const confirmDeleteProduct = (editProduct) => {
    tempData.value = editProduct;
    deleteDialog.value = true;
};

const deleteProduct = () => {
    products.value = products.value.filter((val) => val.id !== tempData.value.id);
    deleteDialog.value = false;
    tempData.value = {};
    toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Deleted', life: 3000 });
};

const findIndexById = (id) => {
    let index = -1;
    for (let i = 0; i < products.value.length; i++) {
        if (products.value[i].ID === id) {
            index = i;
            break;
        }
    }
    return index;
};

const createId = () => {
    let id = '';
    const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    for (let i = 0; i < 5; i++) {
        id += chars.charAt(Math.floor(Math.random() * chars.length));
    }
    return id;
};

const exportCSV = () => {
    dt.value.exportCSV();
};

const confirmDeleteSelected = () => {
    deleteDialogs.value = true;
};
const deleteSelectedProducts = () => {
    products.value = products.value.filter((val) => !selectedProducts.value.includes(val));
    deleteDialogs.value = false;
    selectedProducts.value = null;
    toast.add({ severity: 'success', summary: 'Successful', detail: 'Products Deleted', life: 3000 });
};

const initFilters = () => {
    filters.value = {
        global: { value: null, matchMode: FilterMatchMode.CONTAINS }
    };
};
</script>

<template>
    <div class="grid">
        <div class="col-12">
            <div class="card">
                <Toast />
                <Toolbar class="mb-4">
                    <template v-slot:start>
                        <div class="my-2">
                            <Button label="New" icon="pi pi-plus" class="p-button-success mr-2" @click="openNew" />
                            <Button label="Delete" icon="pi pi-trash" class="p-button-danger" @click="confirmDeleteSelected" :disabled="!selectedProducts || !selectedProducts.length" />
                        </div>
                    </template>

                    <template v-slot:end>
                        <FileUpload mode="basic" accept="image/*" :maxFileSize="1000000" label="Import" chooseLabel="Import" class="mr-2 inline-block" />
                        <Button label="Export" icon="pi pi-upload" class="p-button-help" @click="exportCSV($event)" />
                    </template>
                </Toolbar>

                <DataTable
                    ref="dt"
                    :value="dataList"
                    v-model:selection="selectedProducts"
                    dataKey="id"
                    scrollable
                    :paginator="true"
                    :rows="10"
                    :filters="filters"
                    paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                    :rowsPerPageOptions="[5, 10, 25]"
                    currentPageReportTemplate="Showing {first} to {last} of {totalRecords} products"
                    responsiveLayout="scroll"
                >
                    <template #header>
                        <div class="flex flex-column md:flex-row md:justify-content-between md:align-items-center">
                            <h5 class="m-0">Tools Manage</h5>
                            <span class="block mt-2 md:mt-0 p-input-icon-left">
                                <i class="pi pi-search" />
                                <InputText v-model="filters['global'].value" placeholder="Search..." />
                            </span>
                        </div>
                    </template>

                    <Column selectionMode="multiple" frozen headerStyle="width: 3rem"></Column>
                    <Column field="name" header="name" :sortable="true" frozen></Column>
                    <Column field="type" header="type" :sortable="true"></Column>
                    <Column field="version" header="version" :sortable="true"></Column>
                    <Column field="vendor" header="vendor" :sortable="true"></Column>
                    <Column field="path" header="path" :sortable="true">
                        <template #body="slotProps">
                            <Chip :label="slotProps.data.path" icon="pi pi-copy" cursor-pointer />
                        </template>
                    </Column>
                    <Column header="Operation" alignFrozen="right" frozen headerStyle="min-width:10rem;">
                        <template #body="slotProps">
                            <Button icon="pi pi-pencil" class="p-button-rounded p-button-success mr-2" @click="editProduct(slotProps.data)" />
                            <Button icon="pi pi-trash" class="p-button-rounded p-button-warning mt-2" @click="confirmDeleteProduct(slotProps.data)" />
                        </template>
                    </Column>
                </DataTable>

                <Dialog v-model:visible="editDialog" w-450px header="Tool Details" :modal="true" class="p-fluid">
                    <div class="field">
                        <label for="name">Name</label>
                        <InputText id="name" v-model.trim="tempData.name" required="true" autofocus :class="{ 'p-invalid': submitted && !tempData.name }" />
                        <small color-red class="p-invalid" v-if="submitted && !tempData.name">Name is required.</small>
                    </div>
                    <div class="field">
                        <label for="path">Path</label>
                        <Textarea id="path" v-model="tempData.path" rows="2" cols="20" />
                    </div>
                    <div class="field">
                        <label for="type" class="mb-3">Type</label>
                        <Dropdown id="type" v-model="tempData.type" :options="types" optionLabel="label" optionValue="value" placeholder="Select a Type" required="true" :class="{ 'p-invalid': submitted && !tempData.type }">
                        </Dropdown>
                        <small color-red class="p-invalid" v-if="submitted && !tempData.type">Type is required.</small>
                    </div>
                    <div class="field">
                        <label for="desc">Description</label>
                        <Textarea id="desc" v-model="tempData.desc" required="true" rows="3" cols="20" />
                    </div>
                    <template #footer>
                        <Button label="Cancel" icon="pi pi-times" class="p-button-text" @click="hideDialog" />
                        <Button label="Save" icon="pi pi-check" class="p-button-text" @click="saveData" />
                    </template>
                </Dialog>

                <Dialog v-model:visible="deleteDialog" w-450px header="Confirm" :modal="true">
                    <div class="flex align-items-center justify-content-center">
                        <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
                        <span v-if="tempData">
                            Are you sure you want to delete 
                            <b>{{ tempData.name }}</b>
                            ?
                        </span>
                    </div>
                    <template #footer>
                        <Button label="No" icon="pi pi-times" class="p-button-text" @click="deleteDialog = false" />
                        <Button label="Yes" icon="pi pi-check" class="p-button-text" @click="deleteProduct" />
                    </template>
                </Dialog>

                <Dialog v-model:visible="deleteDialogs" :style="{ width: '450px' }" header="Confirm" :modal="true">
                    <div class="flex align-items-center justify-content-center">
                        <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
                        <span v-if="tempData">Are you sure you want to delete the selected products?</span>
                    </div>
                    <template #footer>
                        <Button label="No" icon="pi pi-times" class="p-button-text" @click="deleteDialogs = false" />
                        <Button label="Yes" icon="pi pi-check" class="p-button-text" @click="deleteSelectedProducts" />
                    </template>
                </Dialog>
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss"></style>
