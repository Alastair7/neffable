import androidx.compose.foundation.layout.*
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController

@Composable
fun MainPage(navController: NavController) {
    Column(
        modifier = Modifier
            .fillMaxSize()
            .padding(16.dp),
        verticalArrangement = Arrangement.SpaceBetween,
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        // Label en la parte superior
        Text(
            text = "NEFFABLE",
            fontSize = 24.sp,
            modifier = Modifier
                .padding(top = 24.dp)
                .align(Alignment.CenterHorizontally)
        )

        Spacer(modifier = Modifier.height(16.dp))

        // Botones en el centro
        Column(
            verticalArrangement = Arrangement.Center,
            horizontalAlignment = Alignment.CenterHorizontally,
            modifier = Modifier.fillMaxHeight()
        ) {
            Button(onClick = { navController.navigate("homeScreen") }) {
                Text(text = "CONNECT")
            }
            Spacer(modifier = Modifier.height(16.dp))
            Button(onClick = { /* Acción para el botón CREATE CONNECTION */ }) {
                Text(text = "CREATE CONNECTION")
            }
        }
    }
}