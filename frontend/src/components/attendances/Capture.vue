<script setup>
import { ref, onMounted } from 'vue';

const inputID = ref('');
const currentAttendee = ref(null);
const selectedEvent = ref(null);
const events = ref([]); // Start as an empty array

// Function to fetch events from your API
const fetchOpenEvents = async () => {
  try {
    // Make sure the URL matches your route group /api/v1/events
    const response = await fetch('http://localhost:8080/api/v1/events?status=open');
    if (!response.ok) throw new Error('Failed to fetch events');

    const results = await response.json();

    // Assuming your API returns { "events": [...] } based on your handler
    events.value = results.data.events || [];
  } catch (error) {
    console.error("Error loading events:", error);
  }
};

const handleScan = async () => {
  // 1. Safety checks
  if (!inputID.value || !selectedEvent.value) return;

  try {
    // 2. Perform the POST request
    const response = await fetch('http://localhost:8080/api/v1/attendances', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        attendee_id: inputID.value,
        event_id: selectedEvent.value.id, // Ensure this matches your Go struct
      }),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to record attendance');
    }

    const result = await response.json();

    console.log(result)

    // 3. Update the UI with response data
    currentAttendee.value = {
      name: result.data.avi.name, // Assuming your API returns attendee details
      public_id: inputID.value,
      school: result.data.avi.school
    };

    // 4. Clear input for next scan
    inputID.value = '';

  } catch (error) {
    console.error("Attendance error:", error);
    alert("Error: " + error.message);
  }
};


onMounted(() => {
  fetchOpenEvents();
});
</script>

<template>
  <v-container max-width="600">
    <h1 class="text-h4 mb-6">Capture Attendance</h1>

    <!-- Dropdown for Event Selection -->
    <v-select
      v-model="selectedEvent"
      :items="events"
      item-title="name"
      item-value="id"
      label="Select Open Event"
      variant="outlined"
      class="mb-4"
      prepend-inner-icon="mdi-calendar-check"
      return-object
    ></v-select>

    <!-- State Display -->
    <v-card class="mb-6" min-height="200">
      <v-card-text class="pa-3">
        <div v-if="!currentAttendee" class="text-center mt-10 text-grey">
          <v-icon size="64" color="grey-lighten-1" class="mb-4">mdi-qrcode-scan</v-icon>
          <p class="text-h6">Waiting for input...</p>
          <p>Please select an event and scan a QR code.</p>
        </div>

        <div v-else>
          <h2 class="text-h5 mb-4 text-primary text-center">Attendee Found!</h2>
          <v-list density="comfortable">
            <v-list-item title="Name" :subtitle="currentAttendee.name"></v-list-item>
            <v-list-item title="ID" :subtitle="currentAttendee.public_id"></v-list-item>
            <v-list-item title="School" :subtitle="currentAttendee.school"></v-list-item>
          </v-list>
        </div>
      </v-card-text>
    </v-card>

    <!-- Input Section -->
    <v-text-field
      v-model="inputID"
      label="Scan or Enter Attendee ID"
      variant="outlined"
      autofocus
      prepend-inner-icon="mdi-barcode-scan"
      @keyup.enter="handleScan"
      :disabled="!selectedEvent"
    ></v-text-field>
  </v-container>
</template>
