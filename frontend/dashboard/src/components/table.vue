<script setup lang="ts">
import { ref, computed } from "vue";
import type { PaginatedStocks } from "../types/interfaces";

const props = defineProps<{
	paginated: PaginatedStocks | null;
}>();

const headers = computed(() => {
	const headers = props.paginated?.stocks?.length ? Object.keys(props.paginated.stocks[0]) : [];
	return headers.filter((header) => header != "id");
});

const formatDate = (dateString: string): string => {
	const date = new Date(dateString);
	return date.toLocaleString("es-MX", {
		weekday: "short",
		year: "numeric",
		month: "long",
		day: "numeric",
		hour: "2-digit",
		minute: "2-digit",
		hour12: false,
	});
};
</script>

<template>
	<div class="overflow-y-auto min-h-60 w-[90vw]">
		<table class="w-full border-collapse" v-if="props.paginated?.stocks">
			<thead class="sticky top-0 bg-white z-10 h-10">
				<tr>
					<th
						v-for="header in headers"
						:key="header"
						class="px-4 py-2 border-b bg-slate-50 text-slate-700 uppercase text-sm"
					>
						{{ header }}
					</th>
				</tr>
			</thead>
			<tbody>
				<tr v-for="(item, index) in props.paginated.stocks" :key="index" class="h-12">
					<td
						v-for="header in headers"
						:key="header"
						class="px-4 py-2 border-b text-slate-600 text-sm"
					>
						{{ header === "time" ? formatDate(item[header]) : item[header] }}
					</td>
				</tr>
			</tbody>
		</table>
		<div v-else>
			<p>No data available</p>
		</div>
	</div>
</template>
