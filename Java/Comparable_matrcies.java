public class Matrix implements Comparable<Matrix> {
    boolean [][]matrix;
    int n, m;
    public Matrix(boolean [][]a){
        this.matrix = a;
        n = a.length;
        m = a[0].length;
    }
    public String toString(){
        String res = "__(matrix) " + String.valueOf(n) + " X " + String.valueOf(m) + " __\n";
        if (n > 3 && m > 3){
            res += (String.valueOf(matrix[0][0]) + " ... " + String.valueOf(matrix[0][m-1]) + "\n ...\n" +
                    String.valueOf(matrix[n-1][0]) + " ... " + String.valueOf(matrix[n-1][m-1]));
        }
        else if (n > 3){
            for (int i = 0; i < m; ++i) res += (String.valueOf(matrix[0][i]) + ' ');
            res += "\n ... \n";
            for (int i = 0; i < m; ++i) res += (String.valueOf(matrix[n-1][i]) + ' ');
        }
        else {
            for (int i = 0; i < n; ++i){
                for (int j = 0; j < m; ++j){
                    res += (String.valueOf(matrix[i][j]) + ' ');
                }
                if (i < n - 1) {res += "\n";}
            }
        }
        return res;
    }
    public int equalsCount(){
        int res = 0;
        res += (n + m);
        for (int i = 0; i < n; ++i){
            for (int j = 0; j < m - 1; ++j){
                if (matrix[i][j] != matrix[i][j+1]){
                    --res;
                    break;
                }
            }
        }
        for (int i = 0; i < m; ++i){
            for (int j = 0; j < n - 1; ++j){
                if (matrix[j][i] != matrix[j+1][i]){
                    --res;
                    break;
                }
            }
        }
        return res;
    }
    public int compareTo(Matrix a){
        if (equalsCount() < a.equalsCount()){
            return -1;
        }
        else if (equalsCount() == a.equalsCount()){
            return 0;
        }
        return 1;
    }

}
