<template>
    <div class="items-center column">
        <div class="justify-between row" style="width: 100%">
            <q-btn color="negative" :disable="row.roleCode === 'super-admin'" @click="handleClear">
                {{ $t('ClearAll') }}
            </q-btn>
            <q-btn color="negative" :disable="row.roleCode === 'super-admin'" @click="handleAll">
                {{ $t('SelectAll') }}
            </q-btn>
            <q-btn color="primary" :disable="row.roleCode === 'super-admin'" @click="handleRoleMenu">
                {{ $t('Save') }}</q-btn>
        </div>
        <q-card-section style="width: 100%; max-height: 70vh" class="scroll">
            <q-tree dense style="width: 100%" :nodes="menuTree" default-expand-all node-key="name" label-key="name"
                selected-color="primary" v-if="menuTree.length !== 0" tick-strategy="strict" v-model:ticked="ticked">
                <template v-slot:default-header="prop">
                    <div class="items-center row">
                        <q-icon :name="prop.node.icon || 'share'" size="28px" class="q-mr-sm" />
                        <div class="text-weight-bold">{{ $t(prop.node.title) }}</div>
                    </div>
                </template>
            </q-tree>
        </q-card-section>
        <q-inner-loading :showing="loading">
            <q-spinner-gears size="50px" color="primary" />
        </q-inner-loading>
    </div>

</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import { ArrayToTree } from 'src/utils/arrayAndTree'
import { postAction, putAction } from 'src/api/manage'

export default {
    name: 'RolePermissionMenu',
    mixins: [tableDataMixin],
    props: {
        row: {
            type: Object,
            required: true,
        },
    },
    computed: {
        menuTree() {
            if (this.tableData.length !== 0) {
                return ArrayToTree(this.tableData, 'name', 'parentCode')
            }
            return []
        },
    },
    data() {
        return {
            pagination: {
                sortBy: 'sort',
                descending: false,
                page: 1,
                rowsPerPage: 10000,
                rowsNumber: 0,
            },
            url: {
                list: 'menu/menu-list',
                roleMenuList: 'role/role-menu',
                roleMenuEdit: 'role/role-menu-edit',
            },
            ticked: [],
        }
    },
    created() {
        this.getTableData()
        this.getRoleMenuList()
    },
    methods: {
        getRoleMenuList() {
            // 每次获取前，清空ticked
            this.ticked = []
            postAction(this.url.roleMenuList, {
                roleCode: this.row.roleCode,
            }).then((res) => {
                if (res.code === 1) {
                    res.data.records.forEach((item) => {
                        this.ticked.push(item.MenuName)
                    })
                }
            })
        },
        handleRoleMenu() {
            const roleMenu = []
            for (let i of this.ticked) {
                roleMenu.push({
                    roleCode: this.row.roleCode,
                    menuName: i,
                })
            }
            putAction(this.url.roleMenuEdit, {
                roleCode: this.row.roleCode,
                roleMenu: roleMenu,
            }).then((res) => {
                if (res.code === 1) {
                    this.$q.notify({
                        type: 'positive',
                        message: res.message,
                    })
                    this.getRoleMenuList()
                }
            })
        },
        handleClear() {
            this.ticked = []
        },
        handleAll() {
            this.tableData.forEach((item) => {
                this.ticked.push(item.name)
            })
        },
    },
}
</script>
