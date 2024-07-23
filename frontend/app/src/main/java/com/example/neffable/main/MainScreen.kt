import androidx.compose.foundation.layout.*
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController

@Composable
fun MainPage(navController: NavController) {
    var newConnectionClicked by remember { mutableStateOf(false) }

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
            if (!newConnectionClicked) {
                Button(
                    onClick = { navController.navigate("homeScreen") },
                    modifier = Modifier.fillMaxWidth().padding(horizontal = 32.dp)
                ) {
                    Text(text = "CONNECT")
                }
                Spacer(modifier = Modifier.height(16.dp))
                Button(
                    onClick = { newConnectionClicked = true },
                    modifier = Modifier.fillMaxWidth().padding(horizontal = 32.dp)
                ) {
                    Text(text = "NEW CONNECTION")
                }
            } else {
                Button(
                    onClick = { /* TODO generar código */ },
                    modifier = Modifier.fillMaxWidth().padding(horizontal = 32.dp)
                ) {
                    Text(text = "GENERATE CODE")
                }
                Spacer(modifier = Modifier.height(16.dp))
                Button(
                    onClick = { /* TODO ingresar código */ },
                    modifier = Modifier.fillMaxWidth().padding(horizontal = 32.dp)
                ) {
                    Text(text = "ENTER CODE")
                }
                Spacer(modifier = Modifier.height(16.dp))
                Button(
                    onClick = { newConnectionClicked = false },
                    modifier = Modifier.padding(horizontal = 32.dp)
                ) {
                    Text(text = "<")
                }
            }
        }
    }
}
