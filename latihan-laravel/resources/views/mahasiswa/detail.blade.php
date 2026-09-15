<h2>Detail Mahasiswa</h2>

<p><strong>NIM:</strong> {{ $mahasiswa->nim }}</p>
<p><strong>Nama:</strong> {{ $mahasiswa->nama }}</p>
<p><strong>Email:</strong> {{ $mahasiswa->email }}</p>

<h3>Daftar Mata Kuliah</h3>

<table border="1">
    <tr>
        <th>Kode</th>
        <th>Nama Mata Kuliah</th>
        <th>SKS</th>
        <th>Semester</th>
        <th>Nilai</th>
    </tr>

    @foreach ($mahasiswa->mataKuliahs as $mataKuliah)
        <tr>
            <td>{{ $mataKuliah->kode }}</td>
            <td>{{ $mataKuliah->nama }}</td>
            <td>{{ $mataKuliah->sks }}</td>
            <td>{{ $mataKuliah->semester }}</td>
            <td>{{ $mataKuliah->pivot->nilai }}</td>
        </tr>
    @endforeach
</table>