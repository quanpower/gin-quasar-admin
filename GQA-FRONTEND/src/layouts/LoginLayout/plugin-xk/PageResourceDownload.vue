<template>
    <div class="q-pa-md row items-center justify-center gqa-resource-download" id="gqa-resource-download">
        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest"
            style="min-width: 40vw; max-width: 45vw">

            <template v-slot:top>
                <span class="text-h6">
                    资源查阅
                </span>
                <q-space />
                <!-- <q-input dense v-model="queryParams.title" label="标题" />
                <q-btn dense color="primary" @click="handleSearch" label="搜索" style="margin: 0 10px" /> -->
                <q-btn dense round push icon="cached" @click="getTableData" />
            </template>

            <template v-slot:body-cell-title="props">
                <q-td :props="props">
                    <q-btn flat dense rounded glossy @click="showDetail(props.row)">
                        {{props.row.title}}
                    </q-btn>
                </q-td>
            </template>

            <template v-slot:body-cell-createdAt="props">
                <q-td :props="props">
                    {{ showDateTime(props.row.createdAt) }}
                </q-td>
            </template>

            <template v-slot:body-cell-createdBy="props">
                <q-td :props="props">
                    <GqaShowName v-if="props.row.createdByUser" :customNameObject="props.row.createdByUser" />
                </q-td>
            </template>

        </q-table>

        <q-dialog v-model="showDetailVisible">
            <q-card style="width: 1400px; max-width: 80vw;">
                <q-card-section class="row items-center q-pb-none">
                    <div class="text-h6">
                        {{detail.title}}
                    </div>
                    <q-space />
                    <q-btn icon="close" flat round dense v-close-popup />
                </q-card-section>

                <q-separator />

                <q-card-section v-html="detail.content" style="max-height: 80vh" class="scroll" />

                <q-separator />

                <q-list bordered separator v-if="detail.attachment">
                    <q-item-label header class="row items-center justify-between">
                        <span>附件列表:</span>
                    </q-item-label>
                    <q-separator />
                    <q-item clickable v-ripple v-for="(item, index) in JSON.parse(detail.attachment)" :key="index"
                        @click="handleDownload(item)">
                        <q-item-section>
                            {{item.filename}}
                        </q-item-section>
                    </q-item>
                </q-list>
            </q-card>
        </q-dialog>
    </div>
</template>

<script>
import { tableDataMixin } from 'src/mixins/tableDataMixin'
import GqaShowName from 'src/components/GqaShowName'
import { FormatDateTime } from 'src/utils/date'
import { downloadAction } from 'src/api/manage'

export default {
    name: 'PageResourceDownload',
    mixins: [tableDataMixin],
    components: {
        GqaShowName,
    },
    computed: {
        showDateTime() {
            return (datetime) => {
                return FormatDateTime(datetime)
            }
        },
    },
    data() {
        return {
            showDetailVisible: false,
            detail: {},
            pagination: {
                sortBy: 'created_at',
                descending: true,
                page: 1,
                rowsPerPage: 10,
            },
            url: {
                list: 'public/plugin-xk/download-list',
            },
            columns: [
                { name: 'title', align: 'center', label: '资源标题', field: 'title' },
                { name: 'createdBy', align: 'center', label: '发布人', field: 'createdBy' },
                { name: 'createdAt', align: 'center', label: '发布时间', field: 'createdAt' },
            ],
        }
    },
    async created() {
        try {
            await this.getTableData()
        } catch (error) {
            this.tableData = [
                {
                    title: 'Gin-Quasar-Admin',
                    content: 'Gin-Quasar-Admin',
                },
            ]
        }
    },
    methods: {
        showDetail(row) {
            this.detail = row
            this.showDetailVisible = true
        },
        handleDownload(item) {
            downloadAction(item.fileUrl.substring(11), item.filename)
        },
    },
}
</script>

<style lang="scss" scoped>
.gqa-resource-download {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>