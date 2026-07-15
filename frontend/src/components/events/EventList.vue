<script setup>
import { ref, onMounted } from 'vue';
import EventDialog from './EventDialog.vue';
import StatusConfirmDialog from './StatusConfirmDialog.vue';
import { useEventStatus } from '../../composables/useEventStatus.js';
import { 
  fetchEvents as apiFetchEvents, 
  saveEvent as apiSaveEvent, 
  updateEventStatus as apiUpdateEventStatus, 
  deleteEvent as apiDeleteEvent 
} from '../../services/eventServices.js';

// --- State ---
const dialog = ref(false);
const loading = ref(false);
const events = ref([]);
const newEvent = ref({ name: '' });

// --- Composable ---
const { statusDialog, itemToChange, statusMessage, statusMap, openStatusDialog } = useEventStatus();

const headers = [
  { title: 'No.', key: 'idx' }, // Uses index slot
  { title: 'Name', key: 'name' },
  { title: 'Status', key: 'status' },
  { title: 'Created at', key: 'created_at' },
  { title: 'Started at', key: 'started_at' },
  { title: 'Action', key: 'action', align: 'center' }
];

// --- API Logic ---
const fetchEvents = async () => {
  loading.value = true;
  try {
    const result = await apiFetchEvents();
    events.value = result.data?.events || result || [];
  } catch (error) {
    console.error('Error fetching events:', error);
  } finally {
    loading.value = false;
  }
};

const saveEvent = async () => {
  try {
    const success = await apiSaveEvent(newEvent.value);
    if (success) {
      await fetchEvents(); // Refresh list
      dialog.value = false;
      newEvent.value = { name: '' };
    }
  } catch (error) {
    console.error('Error saving event:', error);
  }
};

// The action after confirming
const confirmStatusChange = async () => {
  if (!itemToChange.value) return;
  try {
    const nextStatus = itemToChange.value.status === 0 ? 1 : 2; // Example promotion transition
    const success = await apiUpdateEventStatus(itemToChange.value.id, nextStatus);
    if (success) {
      await fetchEvents();
    }
  } catch (error) {
    console.error('Error changing status:', error);
  } finally {
    statusDialog.value = false;
    itemToChange.value = null; // Reset
  }
};

const deleteEvent = async (id) => {
  if (confirm('Are you sure you want to delete this event?')) {
    try {
      const success = await apiDeleteEvent(id);
      if (success) {
        await fetchEvents();
      }
    } catch (error) {
      console.error('Error deleting event:', error);
    }
  }
};

const editEvent = (item) => {
  console.log('Edit event:', item);
};

// --- Lifecycle ---
onMounted(() => {
  fetchEvents();
});
</script>

<template>
  <div class="d-flex justify-space-between align-center mb-4">
    <h1>Event List</h1>
    <v-btn color="primary" @click="dialog = true">Add New Event</v-btn>
  </div>

  <v-data-table
    :headers="headers"
    :items="events"
    :loading="loading"
    class="elevation-1"
  >
    <template v-slot:header.status>
      <div class="text-center">Status</div>
    </template>
    <!-- Incrementing Index -->
    <template v-slot:item.idx="{ index }">
      {{ index + 1 }}
    </template>

    <template v-slot:item.status="{ item }">
      <div class="d-flex justify-center">
        <v-chip
          :color="statusMap[item.status]?.color || 'grey'"
          size="small"
          variant="tonal"
        >
          {{ statusMap[item.status]?.label || 'Unknown' }}
        </v-chip>
      </div>
    </template>

    <!-- Refactored Action Column -->
    <template v-slot:item.action="{ item }">
      <div class="d-flex justify-center gap-2">
        <v-btn
          icon
          variant="text"
          size="small"
          color="primary"
          @click="openStatusDialog(item)"
          :disabled="item.status === 2"
        >
          <v-icon>mdi-cog-refresh</v-icon>
        </v-btn>
        <v-btn
          disabled
          icon
          variant="text"
          size="small"
          color="primary"
          @click="editEvent(item)"
        >
          <v-icon>mdi-pencil</v-icon>
        </v-btn>
        <v-btn
          icon
          variant="text"
          size="small"
          color="error"
          @click="deleteEvent(item.id)"
        >
          <v-icon>mdi-delete</v-icon>
        </v-btn>
      </div>
    </template>
  </v-data-table>

  <!-- Add Event Dialog -->
  <EventDialog
    v-model="dialog"
    :new-event="newEvent"
    @save="saveEvent"
  />

  <!-- Status Confirm Dialog -->
  <StatusConfirmDialog
    v-model="statusDialog"
    :status-message="statusMessage"
    @confirm="confirmStatusChange"
  />
</template>
