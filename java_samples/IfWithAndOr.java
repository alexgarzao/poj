public class IfWithAndOr {
   public static void main(String[] args) {
      int v1 = 111;
      int v2 = 222;
      int v3 = 333;

      System.out.print("v1>v2 and v2>v3? ");
      boolean t = v1 > v2 && v2 > v3;

      if ( t )
         System.out.println(true);
      else
         System.out.println(false);
   }
}
