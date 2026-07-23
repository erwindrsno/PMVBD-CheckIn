<script setup>
  import { ref, onMounted, computed } from 'vue';
  import ViewDetailDialog from './ViewDetailDialog.vue';
  import DeleteConfirmDialog from './DeleteConfirmDialog.vue';

  // --- State ---
  const dialog = ref(false);
  const detailDialog = ref(false); // Add this
  const loading = ref(false);
  const attendees = ref([]);
  const selectedAttendee = ref(null);
  const deleteDialog = ref(false);
  const itemToDelete = ref(null);

  // Form State
  const newAttendee = ref({
    name: '', gender: '', school_id: null, grade_id: null, subgrade_id: null, telp: '', guardian_telp: ''
  });

  // Dropdown Options
  const schools = ref([]);
  const grades = ref([]);
  const subgrades = ref([]);
  const filteredSubgrades = computed(() => {
    // 1. Find the selected grade object
    const selectedGrade = grades.value.find(g => g.id === newAttendee.value.grade_id);

    // 2. If no grade selected or not found, return empty
    if (!selectedGrade) return [];

    // 3. Filter subgrades where subgrade.level_id matches the grade's level_id
    return subgrades.value.filter(s => s.level_id === selectedGrade.level_id);
  });
  const genderOptions = ['Laki-laki', 'Perempuan'];

  const headers = [
    { title: '#', key: 'idx', width: '50px' },
    { title: 'Nama', key: 'name', width: '350px' },
    { title: 'Sekolah', key: 'school'},
    { title: 'Kelas', key: 'grade'},
    { title: 'Telp', key: 'telp' },
    { title: 'Aksi', key: 'action', width:'200px' ,align: 'center' }
  ];

  const viewDetails = (item) => {
    selectedAttendee.value = item;
    detailDialog.value = true;
  };

  const openDeleteDialog = (item) => {
    itemToDelete.value = item;
    deleteDialog.value = true;
  };

  const confirmDelete = async () => {
    if (!itemToDelete.value) return;
    try {
      const token = sessionStorage.getItem('token');
      const response = await fetch(`http://localhost:8080/api/v1/attendees/${itemToDelete.value.public_id}`, {
        method: 'DELETE',
        headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
        },
      });
      if (response.status === 401) {
        sessionStorage.removeItem('token');
        window.location.href = '/login'; // Or use router.push('/login')
        return;
      }
      if (response.ok) {
        await fetchInitialData();
      }
    } catch (error) {
      console.error('Error deleting attendee:', error);
    } finally {
      deleteDialog.value = false;
      itemToDelete.value = null;
    }
  };

  const fetchInitialData = async () => {
    loading.value = true;
    try {
      // Fetch both simultaneously
      const token = sessionStorage.getItem('token');
      const [attendeesRes, schoolsRes, gradesRes, subgradesRes] = await Promise.all([
        fetch('http://localhost:8080/api/v1/attendees', {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
          }
        }),
        fetch('http://localhost:8080/api/v1/schools', {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
          }
        }), // Replace with your actual schools endpoint
        fetch('http://localhost:8080/api/v1/grades', {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
          }
        }), // Replace with your actual schools endpoint
        fetch('http://localhost:8080/api/v1/subgrades', {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
          }
        }) // Replace with your actual schools endpoint
      ]);

      const attendeesData = await attendeesRes.json();
      const schoolsData = await schoolsRes.json();
      const gradesData = await gradesRes.json();
      const subgradesData = await subgradesRes.json();

      attendees.value = attendeesData.data.attendees || [];
      schools.value = schoolsData.data.schools || []; // Assuming your API returns { data: { schools: [...] } }
      grades.value = gradesData.data.grades || [];
      subgrades.value = subgradesData.data.subgrades || [];
    } catch (error) {
      console.error('Error fetching data:', error);
    } finally {
      loading.value = false;
    }
  };

  const saveAttendee = async () => {
    try {
      const token = sessionStorage.getItem('token');
      const response = await fetch('http://localhost:8080/api/v1/attendees', {
        method: 'POST',
        headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify(newAttendee.value),
      });

      if (response.ok) {
        await fetchInitialData(); // Refresh list
        dialog.value = false;
        newAttendee.value = {
          name: '',
          gender: '',
          school_id: null,
          grade_id: null,
          subgrade_id: null,
          telp: '',
          guardian_telp: ''
        };
      }
    } catch (error) {
      console.error('Error saving event:', error);
    }
  };

  onMounted(() => {
    fetchInitialData();
  });
</script>

<template>
  <div class="d-flex justify-space-between align-center mb-4">
    <h1>Daftar peserta</h1>
    <v-btn color="primary" @click="dialog = true">Tambah peserta baru</v-btn>
  </div>

  <v-data-table
    :headers="headers"
    :items="attendees"
    :loading="loading"
    class="elevation-1"
  >
    <!-- Index -->
    <template v-slot:item.idx="{ index }">
      {{ index + 1 }}
    </template>

    <template v-slot:item.grade="{ item }">
      {{ item.grade }} - {{ item.subgrade }}
    </template>

    <!-- Action Column -->
    <template v-slot:header.action>
      <div class="text-center">Aksi</div>
    </template>
    <template v-slot:item.action="{ item }">
      <!-- Call viewDetails(item) instead of editEvent -->
      <v-btn icon variant="text" size="small" color="info" @click="viewDetails(item)">
        <v-icon>mdi-eye</v-icon>
      </v-btn>

      <v-btn disabled icon variant="text" size="small" color="primary" @click="editEvent(item)">
        <v-icon>mdi-pencil</v-icon>
      </v-btn>

      <v-btn icon variant="text" size="small" color="error" @click="openDeleteDialog(item)">
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </template>
  </v-data-table>

  <!-- View Detail Dialog -->
  <ViewDetailDialog v-model="detailDialog" :selected-attendee="selectedAttendee" />

  <!-- Add Attendee Dialog -->
  <v-dialog v-model="dialog" max-width="700px">
    <v-card>
      <v-card-title class="bg-primary text-white">Tambah peserta baru</v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-text-field v-model="newAttendee.name" label="Nama" required></v-text-field>
            </v-col>

            <v-col cols="12">
              <v-radio-group v-model="newAttendee.gender" inline label="Gender">
                <v-radio label="Laki-laki" value="male"></v-radio>
                <v-radio label="Perempuan" value="female"></v-radio>
              </v-radio-group>
            </v-col>

            <v-col cols="12">
              <v-select
                v-model="newAttendee.school_id"
                :items="schools"
                label="Sekolah"
                item-title="name"
                item-value="id"
              ></v-select>
            </v-col>

            <!-- Grade and Subgrade Side by Side -->
            <v-col cols="6">
              <v-select
                v-model="newAttendee.grade_id"
                :items="grades"
                label="Kelas"
                item-title="label"
                item-value="id"
                @update:model-value="newAttendee.subgrade_id = null"
              ></v-select>
            </v-col>
            <v-col cols="6">
              <v-select
                v-model="newAttendee.subgrade_id"
                :items="filteredSubgrades"
                label="Sub kelas"
                item-title="name"
                item-value="id"
                :disabled="!newAttendee.grade_id"
              ></v-select>
            </v-col>

            <!-- Contact numbers Side by Side -->
            <v-col cols="6">
              <v-text-field v-model="newAttendee.telp" label="Nomor telepon"></v-text-field>
            </v-col>
            <v-col cols="6">
              <v-text-field v-model="newAttendee.guardian_telp" label="No. Telp Orang Tua/Wali"></v-text-field>
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="grey" @click="dialog = false">Batal</v-btn>
        <v-btn color="primary" @click="saveAttendee">Simpan</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>


  <DeleteConfirmDialog
    v-model="deleteDialog"
    :attendee-name="itemToDelete?.name || ''"
    @confirm="confirmDelete"
  />
</template>
