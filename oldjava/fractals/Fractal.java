package fractals;

import java.awt.BorderLayout;
import java.awt.Color;
import java.awt.Dimension;
import java.awt.Graphics;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.util.ArrayList;
import java.util.Random;

import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JPanel;

public class Fractal extends JPanel{

	public static final int WINDOW_HEIGHT = 1500;
	public static final int WINDOW_WIDTH = 3000;
	public static final int BUFFER = 10;
	public static final int RANDOM = 100000000;
	public static boolean showAll = true;
	public static boolean doSeed = false;

	
	
	public static void main(String[] args){
		JFrame testFrame = new JFrame();
		testFrame.setDefaultCloseOperation(JFrame.DISPOSE_ON_CLOSE);
		final Fractal f = new Fractal();
		f.setPreferredSize(new Dimension(WINDOW_WIDTH + BUFFER, WINDOW_HEIGHT + BUFFER));
		testFrame.getContentPane().add(f, BorderLayout.CENTER);
		
		JPanel buttonsPanel = new JPanel();
	    JButton doubleButton = new JButton("Double ");
	    JButton clearButton = new JButton("Clear");
	    JButton showButton = new JButton("Show Next");
	    JButton swapButton = new JButton("Swap mode");
	    buttonsPanel.add(doubleButton);
	    buttonsPanel.add(clearButton);
	    buttonsPanel.add(showButton);
	    buttonsPanel.add(swapButton);
	    testFrame.getContentPane().add(buttonsPanel, BorderLayout.SOUTH);
	    doubleButton.addActionListener(new ActionListener() {

	        @Override
	        public void actionPerformed(ActionEvent e) {
	        	f.doubleFractal();
	        }
	    });
	    clearButton.addActionListener(new ActionListener() {

	        @Override
	        public void actionPerformed(ActionEvent e) {
	            f.startOver();
	        }
	    });
	    showButton.addActionListener(new ActionListener() {
	    	@Override
	    	public void actionPerformed(ActionEvent e){
	    		f.showOneMore();
	    	}
	    });
	    swapButton.addActionListener(new ActionListener() {
	    	@Override
	    	public void actionPerformed(ActionEvent e){
	    		showAll = !showAll;
	    	}
	    });
	    testFrame.pack();
	    testFrame.setVisible(true);
		
		//while(true);
	}
	
	@Override
	public void paint(Graphics g){
		super.paintComponent(g);
		g.setColor(Color.black);
		g.fillRect(0, 0, WINDOW_WIDTH + BUFFER, WINDOW_HEIGHT + BUFFER);
		Color c;
		Point p;
		ArrayList<Point> arr = showAll ? points : shown;
		int divisor = info.maxX - info.minX > info.maxY - info.minY ? info.maxX - info.minX : info.maxY - info.minY;
		for(int i = 0; i < arr.size() - 1; i++){
			p = arr.get(i);//points.get(i);
			switch(arr.get(i+1).gen % 7){
			case 0:
				c = Color.red;
				break;
			case 1:
				c = Color.orange;
				break;
			case 2:
				c = Color.yellow;
				break;
			case 3:
				c = Color.green;
				break;
			case 4:
				c = Color.blue;
				break;
			case 5:
				c = Color.magenta;
				break;
			case 6:
				c = Color.pink;
				break;
			default:
				c = Color.black;
			}
			g.setColor(c);
			g.drawLine( ( p.x - info.minX) * WINDOW_WIDTH / divisor + 1, (p.y  - info.minY) * WINDOW_HEIGHT / divisor + 1,
					(arr.get(i + 1).x  - info.minX) * WINDOW_WIDTH / divisor + 1, (arr.get(i + 1).y  - info.minY) * WINDOW_HEIGHT / divisor + 1);		
		}
		/*int i = 0;
		p = points.get(i);
		g.setColor(Color.white);
		while(i < points.size()){
			if(points.get(i).gen <= p.gen) continue;
			g.drawLine( ( p.x - info.minX) * WINDOW_WIDTH / divisor + 1, (p.y  - info.minY) * WINDOW_HEIGHT / divisor + 1,
					(points.get(i).x  - info.minX) * WINDOW_WIDTH / divisor + 1, (points.get(i).y  - info.minY) * WINDOW_HEIGHT / divisor + 1);

		}*/
	}
	
	
	private ArrayList<Point> points;
	private ArrayList<Point> shown;
	private Info info;
	
	public Fractal(){
		startOver();
		doubleFractal();
	}
	
	public void startOver(){
		points = new ArrayList<Point>();
		shown = new ArrayList<Point>();
		points.add(new Point(0, 0, 0));
		points.add(new Point(0, 1, 1));
		shown.add(new Point(0, 0, 0));
		shown.add(new Point(0, 1, 1));
		info = new Info(points.get(0).x,points.get(0).y,points.get(0).x,points.get(0).y);
		doubleFractal();
		repaint();
	}
	
	public ArrayList<Point> getPointsOfFractal(){
		return points;
	}
	
	public void showOneMore(){
		if(shown.size() == points.size())
			doubleFractal();
		shown.add(points.get(shown.size()));
		repaint();
	}
	
	public void doubleFractal(){
		Point axis = points.get(points.size() - 1);
		Random r = doSeed ? new Random(1) : new Random();
		int xoff = 0, yoff = 0;
		for(int i = points.size() - 2; i >= 0; i--){
			int nextX = axis.y - points.get(i).y + axis.x + xoff;
			int nextY = axis.y - axis.x + points.get(i).x + yoff;
			if(r.nextInt(RANDOM) == RANDOM - 1) xoff += r.nextInt(14) - 7;
			if(r.nextInt(RANDOM) == RANDOM - 1) yoff += r.nextInt(14) - 7;
			points.add(new Point(nextX, nextY, axis.gen + 1));
			info.maxX = info.maxX > nextX ? info.maxX : nextX;
			info.minX = info.minX < nextX ? info.minX : nextX;
			info.maxY = info.maxY > nextY ? info.maxY : nextY;
			info.minY = info.minY < nextY ? info.minY : nextY;
		}
		repaint();
	}
	
	public class Point {
		public int x;
		public int y;
		public int gen;
		
		public Point(int gx, int gy, int ggen){
			x = gx;
			y = gy;
			gen = ggen;
		}
	}
	private class Info {
		public int minX;
		public int minY;
		public int maxX;
		public int maxY;
		
		public Info(int x1, int y1, int x2, int y2){
			minX = x1;
			minY = y1;
			maxX = x2;
			maxY = y2;
		}
	}
}