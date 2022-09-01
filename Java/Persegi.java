import java.lang.Math;

public class Persegi {
    int sisi = 5;

    public int luas () {
        int luas =(int) Math.pow(this.sisi, 2);
        return luas;
    }

    public double volume () {
        double volume = Math.pow(Double.valueOf(this.sisi), 3);
        return volume;
    }

    public static void main(String[] args) {
        Persegi sd = new Persegi();
        System.out.println("ini try");
        System.out.println("sisi dari perseginya adalah : "+sd.sisi);
        System.out.println("luas dari perseginya adalah : "+sd.luas());
        System.out.println("volume dari perseginya adalah : "+sd.volume());
    }
}


